package link

import "time"

type Link struct {
	Code          string
	Orig_url      string `json:"original_url"`
	Owner_id      int64  `json:"-"`
	Monet_enabled bool   `json:"monetization_enabled"`
	Expires_at    *time.Time
	Duration      int `json:"duration"`
}

type PatchLink struct {
	Code          string
	Orig_url      *string `json:"original_url"`
	Owner_id      *int64  `json:"-"`
	Monet_enabled bool    `json:"monetization_enabled"`
	Expires_at    *time.Time
	Duration      int `json:"duration"`
}

type BindLink struct {
	Code          string
	Orig_url      string `json:"original_url"`
	Monet_enabled bool   `json:"monetization_enabled"`
	Duration      int    `json:"duration"`
}
