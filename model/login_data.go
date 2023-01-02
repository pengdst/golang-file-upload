package model

type LoginData struct {
	User  User  `json:"user"`
	Token Token `json:"token"`
}
