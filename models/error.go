package models

type Error struct {
	Message    string `json:"message"`
	StatusCode int64  `json:"status_code"`
}
