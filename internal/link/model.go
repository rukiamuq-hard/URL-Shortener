package link

import "time"

type Link struct {
	Code       string
	Orig_url   string `json:"original_url"`
	Created_at time.Time
	Expires_at time.Time
	Duration   int `json:"duration"` // time what link live
}
