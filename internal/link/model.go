package link

import "time"

type Link struct {
	Code       string `json:"code"`
	Orig_url   string `json:"original_url"`
	Created_at time.Time
	Duration   time.Time `json:"duration"` // time what link live
}
