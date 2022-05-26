package services

import (
	"os"
	"time"
	"unisun/api/auth-listener/src/constants"
	"unisun/api/auth-listener/src/logging"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
}

type jwtServices struct {
	SecretKey string `json:"secretKey"`
	Issure    string `json:"issure"`
}

func JWTAuthService() JWTService {
	return &jwtServices{
		SecretKey: getSecretKey(),
		Issure:    "Bikash",
	}
}

func getSecretKey() string {
	secret := os.Getenv(constants.JWT_SECRET)
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func GenerateRefreshJWT(token_version int, user_id int) (string, error) {
	var mySigningKey = []byte(os.Getenv(constants.JWT_SECRET))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["token_version"] = token_version
	claims["iat"] = time.Now().Unix()
	claims["uid"] = user_id
	claims["exp"] = time.Now().Add(86400 * time.Minute).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		logging.Println("Something Went Wrong: ", err.Error())
		return "", err
	}
	return tokenString, nil
}
