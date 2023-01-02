package model

import "time"

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
