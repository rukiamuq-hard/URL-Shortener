package link

import "time"

type Link struct {
	Code       string `json:"code"`
	Orig_url   string `json:"original_url"`
	Created_at time.Time
	Expired_at time.Time
	Duration   int `json:"duration"` // time what link live
}
