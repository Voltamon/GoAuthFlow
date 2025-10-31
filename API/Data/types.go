package data

import "time"

type User struct {
	ID        int       `json:"id"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
}
