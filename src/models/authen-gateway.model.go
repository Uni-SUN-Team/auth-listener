package models

type CallAuthenGatewayResponse struct {
	Error  string            `json:"error"`
	Result map[string]string `json:"result"`
}
