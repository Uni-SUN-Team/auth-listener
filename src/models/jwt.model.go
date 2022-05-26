package models

type RefreshJWT struct {
	TokenVersion int   `json:"token_version"`
	Uid          int   `json:"uid"`
	Iat          int64 `json:"iat"`
	Ext          int64 `json:"ext"`
}

type TokenJWT struct {
	Id  int   `json:"id"`
	Iat int64 `json:"iat"`
	Ext int64 `json:"ext"`
}
