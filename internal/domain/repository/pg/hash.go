package pg

import (
	"context"
	"github.com/Akanibekuly/tsarka-tz/internal/domain/entities"
	"github.com/Akanibekuly/tsarka-tz/internal/domain/errs"
	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
	"github.com/jackc/pgx/v4"
)

type HashRepository struct {
	lg   interfaces.Logger
	conn *pgx.Conn
}

func NewHashRepository(lg interfaces.Logger, conn *pgx.Conn) *HashRepository {
	return &HashRepository{
		lg:   lg,
		conn: conn,
	}
}

func (d *HashRepository) HashCreate(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := d.conn.Exec(ctx, `INSERT INTO hash (id) VALUES ($1)`, id)
	if err != nil {
		d.lg.Errorw("[DATABASE] hash create", err, "id", id)
		return err
	}

	return nil
}

func (d *HashRepository) HashGet(id string) (*entities.HashSt, error) {
	var hash entities.HashSt

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := d.conn.QueryRow(ctx, `SELECT status, result FROM hash WHERE id=$1`, id).
		Scan(
			&hash.Status,
			&hash.Result,
		)
	if err != nil {
		d.lg.Errorw("[DATABASE] hash get", err, "id", id)
		if err == pgx.ErrNoRows {
			return nil, errs.ObjectNotFound
		}
		return nil, errs.InternalServerError
	}

	return &hash, nil
}

func (d *HashRepository) HashUpdate(id string, hash *entities.HashSt) error {
	if hash == nil {
		return errs.PointerIsNil
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := "UPDATE hash SET status=$1, result=$2 WHERE id=$3"
	_, err := d.conn.Exec(ctx, query, hash.Status, hash.Result, id)
	if err != nil {
		d.lg.Errorw("[DATABASE] hash update", err, "hash", hash, "id", id)
		return errs.InternalServerError
	}

	return nil
}
