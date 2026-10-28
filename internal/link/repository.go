package link

import (
	"context"
	"database/sql"
)

type Postgres interface {
	CreateLink(ctx context.Context, l *Link) error
}

type Repository struct {
	Db *sql.DB
}

func NewRepository(Db *sql.DB) *Repository {
	return &Repository{
		Db: Db,
	}
}

func (r *Repository) CreateLink(ctx context.Context, l *Link) error {
	_, err := r.Db.ExecContext(ctx,
		`INSERT INTO links(code, original_url, monetization_enabled, expires_at) VALUES($1, $2, $3, $4)`,
		l.Code, l.Orig_url, l.Monet_enabled, l.Expires_at,
	)
	return err
}
