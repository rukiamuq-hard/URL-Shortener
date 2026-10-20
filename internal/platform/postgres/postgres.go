package postgres

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Postgres struct {
	DB *sql.DB
}

func New() *Postgres {
	return &Postgres{}
}

func (pg *Postgres) Connect(postgresDSN string, ctx context.Context) error {
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var err error
	pg.DB, err = sql.Open("pgx", postgresDSN)
	if err != nil {
		return err
	}

	if err = pg.DB.PingContext(pingCtx); err != nil {
		pg.Close()
		return err
	}

	return nil
}

func (pg *Postgres) Close() error {
	if err := pg.DB.Close(); err != nil {
		return err
	}
	return nil
}
