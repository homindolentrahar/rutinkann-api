package web

import "com.homindolentrahar.rutinkann-api/model"

type AuthResponse struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}
