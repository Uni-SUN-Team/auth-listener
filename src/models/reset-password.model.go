package models

type ResetPasswordRequest struct {
	Code                 string `json:"code"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type ResetPasswordResponse struct {
	Data  ResetPasswordResponseDetail `josn:"data"`
	Error ResetPasswordResponseError  `json:"error"`
}

type ResetPasswordResponseDetail struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ResetPasswordResponseError struct {
	Status  int    `json:"status"`
	Name    string `json:"name"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
