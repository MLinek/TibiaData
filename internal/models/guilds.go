package models

import "time"

type Guilds struct {
	Guilds struct {
		Guild Guild `json:"guild"`
	} `json:"guilds"`
	Information struct {
		ApiVersion int       `json:"api_version"`
		Timestamp  time.Time `json:"timestamp"`
	} `json:"information"`
}
