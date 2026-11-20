package link

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Repository interface {
	CreateLink(ctx context.Context, l *Link) error
	GetByCode(ctx context.Context, code string) (*Link, error)
}

type PostgresRepository struct {
	Db *sql.DB
}

func NewPostgresRepository(Db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		Db: Db,
	}
}

func (r *PostgresRepository) CreateLink(ctx context.Context, l *Link) error {
	_, err := r.Db.ExecContext(ctx,
		`INSERT INTO links(code, original_url, monetization_enabled, expires_at) VALUES($1, $2, $3, $4)`,
		l.Code, l.Orig_url, l.Monet_enabled, l.Expires_at,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "links_code_key" {
			return ErrCodeTaken
		}
	}
	return err
}

func (r *PostgresRepository) GetByCode(ctx context.Context, code string) (*Link, error) {
	var myLink Link
	err := r.Db.QueryRowContext(ctx,
		`SELECT code, original_url, monetization_enabled, expires_at FROM links WHERE code = $1`,
		code,
	).Scan(&myLink.Code, &myLink.Orig_url, &myLink.Monet_enabled, &myLink.Expires_at)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &myLink, nil
}
