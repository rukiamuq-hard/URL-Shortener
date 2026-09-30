package postgres

import (
	"database/sql"
)

type Postgress struct {
	db *sql.DB
}

func New() *Postgress {
	return &Postgress{
		db: &sql.DB{},
	}
}
