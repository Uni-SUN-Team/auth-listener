package models

type SigninCallRequest struct {
	Token  string `json:"refresh_token"`
	UserId int    `json:"user_id"`
}
