package models

type ServiceIncomeRequest struct {
	Path   string `json:"path"`
	Method string `json:"method"`
	Body   []byte `json:"body"`
}

type ServiceIncomeResponse struct {
	Error   string `json:"error"`
	Status  bool   `json:"status"`
	Payload string `json:"payload"`
}
