package models

import "time"

type UserAuthPermission struct {
	UserId       int
	TokenVersion int
	Iat          float64
	Ext          float64
	LastLogin    time.Time
}
