package models

type RefreshTokenBodyRequest struct {
	Token string `json:"token"`
}

type RefreshTokenBodyResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Claims  RefreshJWT `json:"claims"`
}

type RefreshTokenResponse struct {
	Token        string `josn:"token"`
	RefreshToken string `json:"refresh_token"`
}
