package models

import "time"

type SigninPayload struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type Signin struct {
	Jwt        string `json:"jwt"`
	JwtRefresh string `json:"jwt_refresh"`
	User       User   `json:"user"`
}

type User struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Provider   string    `json:"provider"`
	Confirmed  bool      `json:"confirmed"`
	Blocked    bool      `json:"blocked"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"lastname"`
	Dob        string    `json:"dob"`
	Gender     string    `json:"gender"`
	Telephone  string    `json:"telephone"`
	Agreement  bool      `json:"agreement"`
	Middlename string    `json:"middlename"`
}

type SignWarning struct {
	Data  map[string]string `json:"data"`
	Error SignWarningError  `json:"error"`
}

type SignWarningError struct {
	Status  int         `json:"status"`
	Name    string      `json:"name"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
	Error   error       `json:"error"`
}
