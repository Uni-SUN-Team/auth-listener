package models

type ResponseFail struct {
	Data  string      `josn:"data"`
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Status  int    `json:"status"`
	Name    string `json:"name"`
	Message string `json:"message"`
	Detail  error  `json:"detail"`
}
