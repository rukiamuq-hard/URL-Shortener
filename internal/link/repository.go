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
	DeleteExpired(ctx context.Context) (int64, error)
	DeleteLinks(ctx context.Context, code string) error
	UpdateLink(ctx context.Context, l *PatchLink) error
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

func (r *PostgresRepository) DeleteExpired(ctx context.Context) (int64, error) {
	res, err := r.Db.ExecContext(ctx, `DELETE FROM links WHERE expires_at IS NOT NULL AND expires_at < NOW()`)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *PostgresRepository) DeleteLinks(ctx context.Context, code string) error {
	_, err := r.Db.ExecContext(ctx, `DELETE FROM links WHERE code = $1`, code)
	return err
}

func (r *PostgresRepository) UpdateLink(ctx context.Context, l *PatchLink) error {
	_, err := r.Db.ExecContext(ctx,
		`UPDATE links SET monetization_enabled = COALESCE($2, monetization_enabled),
		expires_at = COALESCE($3, expires_at),
		original_url = COALESCE($4, original_url)
    WHERE code = $1
		`, l.Code, l.Monet_enabled, l.Expires_at, l.Orig_url)
	return err
}
