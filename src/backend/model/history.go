package model

import "time"

type History struct {
	Content    string    `json:"content"`
	FromUser   bool      `json:"fromUser"`
	Created_At time.Time `json:"created_at"`
}