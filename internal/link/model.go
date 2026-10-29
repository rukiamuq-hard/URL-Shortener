package link

import "time"

type Link struct {
	Code     string
	Orig_url string `json:"original_url"`
	//	Owner_id      int64
	Monet_enabled bool `json:"monetization"`
	Expires_at    *time.Time
	Duration      int `json:"duration"` // time what link live
}
