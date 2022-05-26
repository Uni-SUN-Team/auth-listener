package models

type ResponseRefreshStrapi struct {
	Jwt string `json:"jwt"`
}

type StrapiError struct {
	Data  interface{}       `json:"data"`
	Error StrapiErrorDetail `json:"error"`
}

type StrapiErrorDetail struct {
	Status  string      `json:"status"`
	Name    string      `json:"name"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}
