package models

type ForgetPasswordRequest struct {
	Email string `json:"email"`
}

type ForgetPasswordResponse struct {
	Data  ForgetPasswordResponseDetail `josn:"data"`
	Error ForgetPasswordResponseError  `json:"error"`
}

type ForgetPasswordResponseDetail struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ForgetPasswordResponseError struct {
	Status  int    `json:"status"`
	Name    string `json:"name"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
