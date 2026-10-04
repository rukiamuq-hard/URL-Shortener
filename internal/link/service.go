package link

import (
	"URLS/internal/monetization"
)

type Service struct {
	AdServ *monetization.AdProvider
}

func NewService(AdServ *monetization.AdProvider) *Service {
	return &Service{
		AdServ: AdServ,
	}
}
