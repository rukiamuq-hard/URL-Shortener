package link

import (
	"context"
	"time"
)

type Service struct {
	PDB Postgres
}

func NewService(PDB Postgres) *Service {
	return &Service{
		PDB: PDB,
	}
}

func (s *Service) CreateLink(ctx context.Context, l *Link) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	l.Created_at = time.Now()
	l.Expired_at = l.Created_at.Add(time.Duration(l.Duration) * time.Second)

	if err := s.PDB.CreateLink(ctx, l); err != nil {
		return err
	}
	return nil
}
