package pg

import (
	"context"
	"fmt"
	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
	"time"

	"github.com/Akanibekuly/tsarka-tz/internal/domain/entities"
	"github.com/Akanibekuly/tsarka-tz/internal/domain/errs"
	"github.com/jackc/pgx/v4"
)

const timeout = time.Second * 5

type UserRepository struct {
	lg   interfaces.Logger
	conn *pgx.Conn
}

func NewUserRepository(lg interfaces.Logger, conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		lg:   lg,
		conn: conn,
	}
}

func (d *UserRepository) UserGet(id int) (*entities.UserSt, error) {
	query := `SELECT 
				first_name, last_name
				FROM users WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var user entities.UserSt
	err := d.conn.QueryRow(ctx, query, id).Scan(
		&user.FirstName,
		&user.LastName,
	)
	if err != nil {
		d.lg.Errorw("[DATABASE] get user", err)
		if err == pgx.ErrNoRows {
			return nil, errs.ObjectNotFound
		}
		return nil, errs.InternalServerError
	}

	return &user, nil
}

func (d *UserRepository) UserCreate(user *entities.UserSt) (int, error) {
	if user == nil {
		return 0, errs.PointerIsNil
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := `INSERT INTO users (first_name, last_name)
				VALUES ($1,$2) RETURNING id`

	var id int
	err := d.conn.QueryRow(ctx, query, user.FirstName, user.LastName).Scan(&id)
	if err != nil {
		d.lg.Errorw("[DATABASE] create user", err, "user", user)
		return 0, errs.InternalServerError
	}

	return id, nil
}

func (d *UserRepository) UserUpdate(id int, user *entities.UserUpdateSt) error {
	if user == nil {
		return errs.PointerIsNil
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := `UPDATE users SET`
	var args []interface{}
	if user.FirstName != nil {
		args = append(args, *user.FirstName)
		query += fmt.Sprintf(" first_name=$%d", len(args))
	}
	if user.LastName != nil {
		args = append(args, *user.LastName)
		if len(args) == 2 {
			query += ","
		}
		query += fmt.Sprintf(" last_name=$%d", len(args))
	}

	args = append(args, id)
	query += fmt.Sprintf(` WHERE id = $%d`, len(args))

	_, err := d.conn.Exec(ctx, query, args...)
	if err != nil {
		d.lg.Errorw("[DATABASE] update user", err, "user", user, "id", id, "query", query)
		return errs.InternalServerError
	}

	return nil
}

func (d *UserRepository) UserDelete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := d.conn.Exec(ctx, `DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		d.lg.Errorw("[DATABASE] delete user", err, "id", id)
		return errs.InternalServerError
	}

	return nil
}
