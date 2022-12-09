package model

import "time"

type User struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Token      string    `json:"token"`
	CreatedAt  time.Time `json:"created_at"`
	UpdaetedAt time.Time `json:"updaeted_at"`
}
