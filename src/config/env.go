package config

import (
	"os"
	"unisun/api/auth-listener/src/constants"
)

func SetENV() {
	os.Setenv(constants.JWT_SECRET, "aSiAZgPRmmw7gN7p9WeQxQ==")
	os.Setenv(constants.CONTEXT_PATH, "/auth")
	/**
	* Highlights: Strapi information gateway
	 */
	os.Setenv(constants.HOST_STRAPI_SERVICE, "http://localhost:8082")
	os.Setenv(constants.PATH_STRAPI_INFORMATION_GATEWAY, "/strapi-information-gateway/api/strapi")
	os.Setenv(constants.PATH_STRAPI_SIGNIN, "/api/auth/local")
	os.Setenv(constants.PATH_STRAPI_REFRESHTOKEN, "/api/auth/refreshToken")
	os.Setenv(constants.PATH_STRAPI_FORGET_PASSWORD, "/api/auth/forgot-password")
	os.Setenv(constants.PATH_STRAPI_RESET_PASSWORD, "/api/auth/reset-password")
	os.Setenv(constants.PATH_STRAPI_REGISTER, "/api/auth/local/register")
	/**
	* Highlights: Authen gateway
	 */
	os.Setenv(constants.AUTHEN_GATEWAY_HOST, "http://localhost:8081")
	os.Setenv(constants.AUTHEN_GATEWAY_PATH_SIGNIN, "/authen-listening/api/validate/call-signin")
	os.Setenv(constants.AUTHEN_GATEWAY_PATH_GET_TOKENVERSION, "/authen-listening/api/validate/token-version/")
	os.Setenv(constants.AUTHEN_GATEWAY_PATH_CALL_REVOKE, "/authen-listening/api/validate/call-revoke")
	os.Setenv(constants.AUTHEN_GATEWAY_PATH_CALL_REFRESHTOKEN, "/authen-listening/api/validate/call-check-refreshtoken")

	os.Setenv(constants.LOG_PATH, "/Users/ns/Documents/UniSUN/auth-listener/tmp/app.log")
	os.Setenv(constants.PATH_STRAPI_CALLBACK_GOOGLE, "/api/auth/google/callback")
	os.Setenv(constants.PATH_STRAPI_CALLBACK_FACEBOOK, "/api/auth/facebook/callback")
}
