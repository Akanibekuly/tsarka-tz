package pg

import (
	"context"

	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
	"github.com/jackc/pgx/v4"
)

type St struct {
	lg   interfaces.Logger
	conn *pgx.Conn
}

func New(lg interfaces.Logger) *St {
	return &St{
		lg: lg,
	}
}

func (d *St) Connect(dsn string) error {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		d.lg.Errorw("[DATABASE] connect", err)
		return err
	}

	d.conn = conn

	return nil
}

func (d *St) Close() {
	if err := d.conn.Close(context.Background()); err != nil {
		d.lg.Errorw("[DATABASE] close conn", err)
	}
}
