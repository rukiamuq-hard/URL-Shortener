package link

import (
	"URLS/pkg/shortcode"
	"context"
	"errors"
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
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	var err error

	u, err := url.Parse(l.Orig_url)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return nil, ErrInvalidInput
	}

	if l.Duration < 0 {
		return nil, ErrExpired
	} else if l.Duration > 0 {
		t := time.Now().Add(time.Duration(l.Duration) * time.Second)
		l.Expires_at = &t
	}

	for i := 0; i < 5; i++ {
		l.Code, err = s.sc.GenerateCode(7)
		if err != nil {
			return nil, err
		}

		err := s.pdb.CreateLink(ctx, l)
		if !errors.Is(err, ErrCodeTaken) {
			return l, err
		}
	}

	return nil, ErrCodeTaken
}
