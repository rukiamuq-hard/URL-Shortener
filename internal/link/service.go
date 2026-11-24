package link

import (
	"URLS/pkg/shortcode"
	"context"
	"errors"
	"log/slog"
	"net/url"
	"time"
)

type Service struct {
	pdb Repository
	sc  *shortcode.Shortcode
}

func NewService(PDB Repository, SC *shortcode.Shortcode) *Service {
	return &Service{
		pdb: PDB,
		sc:  SC,
	}
}

func (s *Service) CreateLink(ctx context.Context, l *Link) (*Link, error) {
	dbctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var err error

	u, err := url.Parse(l.Orig_url)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return nil, ErrInvalidInput
	}

	if l.Duration < 0 {
		return nil, ErrInvalidInput
	} else if l.Duration > 0 {
		t := time.Now().Add(time.Duration(l.Duration) * time.Second)
		l.Expires_at = &t
	}

	for i := 0; i < 5; i++ {
		l.Code, err = s.sc.GenerateCode(7)
		if err != nil {
			return nil, err
		}

		err := s.pdb.CreateLink(dbctx, l)
		if !errors.Is(err, ErrCodeTaken) {
			return l, err
		}
	}

	return nil, ErrCodeTaken
}

func (s *Service) GetLinkInfo(ctx context.Context, code string) (*Link, error) {
	dbctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	val, err := s.pdb.GetByCode(dbctx, code)
	if err != nil {
		return nil, err
	}

	if val.Expires_at != nil && time.Now().After(*val.Expires_at) {
		return nil, ErrExpired
	}

	return val, nil
}

func (s *Service) clean(ctx context.Context) error {
	clctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	n, err := s.pdb.DeleteExpired(clctx)
	if err != nil {
		return err
	}
	if n > 0 {
		slog.Info("deleted", "nums", n)
	}
	return nil
}

func (s *Service) Cleaner(ctx context.Context, every time.Duration) {
	err := s.clean(ctx)
	if err != nil {
		slog.Error("error cleanup ", "err", err.Error())
	}
	timer := time.NewTicker(1 * every)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			err := s.clean(ctx)
			if err != nil {
				slog.Error("error cleanup ", "err", err.Error())
			}
		case <-ctx.Done():
			slog.Info("Cleaner stoped by context.")
			return
		}
	}

}

func (s *Service) Delete(ctx context.Context, code string) error {
	sctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := s.pdb.DeleteLinks(sctx, code)
	return err
}

func (s *Service) Update(ctx context.Context, l *PatchLink) error {
	sctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if l.Orig_url != nil {
		u, err := url.Parse(*l.Orig_url)
		if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Hostname() == "" {
			return ErrInvalidInput
		}
	}

	if l.Duration < 0 {
		return ErrInvalidInput
	} else if l.Duration > 0 {
		t := time.Now().Add(time.Duration(l.Duration) * time.Second)
		l.Expires_at = &t
	}

	err := s.pdb.UpdateLink(sctx, l)
	return err
}
