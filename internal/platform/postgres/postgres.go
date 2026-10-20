package postgres

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Postgress struct {
	DB *sql.DB
}

func New() *Postgress {
	return &Postgress{}
}

func (pg *Postgress) Connect(postgresDSN string) error {
	var err error
	pg.DB, err = sql.Open("pgx", postgresDSN)
	if err != nil {
		return err
	}
	defer pg.DB.Close()

	if err = pg.DB.Ping(); err != nil {
		return err
	}

	return nil
}

func (pg *Postgress) Close() error {
	if err := pg.DB.Close(); err != nil {
		return err
	}
	return nil
}
