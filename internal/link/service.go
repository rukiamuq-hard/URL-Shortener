package link

import (
	"context"
	"crypto/rand"
	"math/big"
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

	var err error

	l.Code, err = s.generateCode()
	if err != nil {
		return err
	}

	l.Created_at = time.Now()
	l.Expires_at = l.Created_at.Add(time.Duration(l.Duration) * time.Second)

	if err := s.PDB.CreateLink(ctx, l); err != nil {
		return err
	}
	return nil
}

func (s *Service) generateCode() (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const codeLen = 7

	result := make([]byte, codeLen)
	for i := 0; i < codeLen; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result[i] = letters[num.Int64()]
	}
	return string(result), nil
}
