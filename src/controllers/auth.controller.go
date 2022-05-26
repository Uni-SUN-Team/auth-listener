package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"unisun/api/auth-listener/src/constants"
	"unisun/api/auth-listener/src/logging"
	"unisun/api/auth-listener/src/models"
	"unisun/api/auth-listener/src/services"

	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context) {
	payloadRequestSignin := models.ServiceIncomeRequest{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Println("", err.Error())
	}
	payloadRequestSignin.Path = os.Getenv(constants.PATH_STRAPI_SIGNIN)
	payloadRequestSignin.Method = constants.POST
	payloadRequestSignin.Body = jsonData
	payloadSignin := models.Signin{}
	data := services.GetInformationFormStrapi(payloadRequestSignin)
	err = json.Unmarshal([]byte(data.Payload), &payloadSignin)
	if err != nil {
		logging.Println("Change byte to json", err.Error())
		c.JSON(http.StatusUnprocessableEntity, data)
		return
	}
	if payloadSignin.Jwt != "" && payloadSignin.User.Confirmed && payloadSignin.User.Email != "" {
		signinCallRequest := models.SigninCallRequest{}
		if userPermission, err := services.GetUserPermission(payloadSignin.User.Id); err != nil {
			logging.Println("Call get user-auth-permission is Error.", err.Error())
		} else {
			if userPermission.UserId != 0 && userPermission.TokenVersion != 0 {
				if jwt, err := services.GenerateRefreshJWT(userPermission.TokenVersion+1, userPermission.UserId); err != nil {
					logging.Println("Generate refresh token Error.", err.Error())
					c.JSON(http.StatusBadRequest, &models.ResponseFail{Error: models.ErrorDetail{
						Detail:  err,
						Name:    "AppError",
						Message: err.Error(),
						Status:  http.StatusBadRequest,
					}})
				} else {
					payloadSignin.JwtRefresh = jwt
					signinCallRequest.Token = jwt
					signinCallRequest.UserId = userPermission.UserId
				}
			} else {
				if jwt, err := services.GenerateRefreshJWT(1, payloadSignin.User.Id); err != nil {
					logging.Println("Generate refresh token Error.", err.Error())
					c.JSON(http.StatusBadRequest, &models.ResponseFail{Error: models.ErrorDetail{
						Detail:  err,
						Name:    "AppError",
						Message: err.Error(),
						Status:  http.StatusBadRequest,
					}})
				} else {
					payloadSignin.JwtRefresh = jwt
					signinCallRequest.Token = jwt
					signinCallRequest.UserId = payloadSignin.User.Id
				}
			}
		}
		if responseCallSignin, err := services.CallSignIn(signinCallRequest); err != nil {
			c.JSON(http.StatusBadRequest, &models.ResponseFail{Error: models.ErrorDetail{
				Detail:  err,
				Name:    "AppError",
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			}})
			logging.Println("Call signin is Error.", err.Error())
		} else {
			if responseCallSignin.Result["confirm"] == "true" {
				c.JSON(http.StatusOK, payloadSignin)
			} else {
				c.JSON(http.StatusBadRequest, &models.ResponseFail{Error: models.ErrorDetail{
					Detail:  nil,
					Name:    "CallError",
					Message: "confirm=" + responseCallSignin.Result["confirm"] + responseCallSignin.Error,
					Status:  http.StatusBadRequest,
				}})
			}
		}
	} else {
		signWarning := models.SignWarning{}
		err = json.Unmarshal([]byte(data.Payload), &signWarning)
		if err != nil {
			logging.Println("Change byte to json", err.Error())
			c.JSON(http.StatusUnprocessableEntity, data)
			return
		}
		c.JSON(http.StatusFound, signWarning)
	}
}

func Signout(c *gin.Context) {
	payload := models.Revoke{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Println("", err.Error())
	}
	err = json.Unmarshal([]byte(jsonData), &payload)
	if err != nil {
		logging.Println("Change byte to json", err.Error())
		return
	} else {
		err = nil
	}
	if responseCall, err := services.CallRevoke(payload); err != nil {
		logging.Println("Call revoke is Error.", err.Error())
		c.JSON(http.StatusBadRequest, responseCall)
	} else {
		c.JSON(http.StatusOK, responseCall.Result)
	}
}

func RefreshToken(c *gin.Context) {
	body := models.RefreshTokenBodyRequest{}
	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Println("", err.Error())
	}
	err = json.Unmarshal([]byte(requestBody), &body)
	if err != nil {
		logging.Println("Change byte to json", err.Error())
		return
	}
	responseRefreshToken, err := services.CallRefreshToken(body)
	if err != nil {
		logging.Println("Call api refresh token is Error.", err.Error())
		return
	}
	token := models.ResponseRefreshStrapi{}
	if responseRefreshToken.Status {
		payloadRequest := models.ServiceIncomeRequest{}
		payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_REFRESHTOKEN)
		payloadRequest.Method = constants.POST
		payload, err := json.Marshal(body)
		if err != nil {
			logging.Println("Convert json to []byte is error.", err.Error())
			c.JSON(http.StatusUnprocessableEntity, &models.RefreshTokenResponse{})
			return
		}
		payloadRequest.Body = payload
		data := services.GetInformationFormStrapi(payloadRequest)
		err = json.Unmarshal([]byte(data.Payload), &token)
		if err != nil {
			logging.Println("Change byte to json.", err.Error())
			c.JSON(http.StatusUnprocessableEntity, &models.RefreshTokenResponse{})
			return
		}
		jwt_refresh, err := services.GenerateRefreshJWT(responseRefreshToken.Claims.TokenVersion, responseRefreshToken.Claims.Uid)
		if err != nil {
			logging.Println("Generate refresh token Error.", err.Error())
			c.JSON(http.StatusBadRequest, &models.RefreshTokenResponse{})
			return
		}
		c.JSON(http.StatusOK, &models.RefreshTokenResponse{
			Token:        token.Jwt,
			RefreshToken: jwt_refresh,
		})
	} else {
		c.JSON(http.StatusFound, &models.RefreshTokenResponse{})
		return
	}
}

func ForgetPassword(c *gin.Context) {
	requestStrapi := models.ServiceIncomeRequest{}
	jsonRequest, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Println("Request is error.", err.Error())
	}
	requestStrapi.Path = os.Getenv(constants.PATH_STRAPI_FORGET_PASSWORD)
	requestStrapi.Method = constants.POST
	requestStrapi.Body = jsonRequest
	responseStrapi := services.GetInformationFormStrapi(requestStrapi)
	if responseStrapi.Status {
		c.JSON(http.StatusOK, &models.ForgetPasswordResponse{
			Data: models.ForgetPasswordResponseDetail{
				Status:  responseStrapi.Status,
				Message: "Your user received an email.",
			},
		})
	} else {
		strapiResponse := models.StrapiError{}
		if err := json.Unmarshal([]byte(responseStrapi.Payload), &strapiResponse); err != nil {
			logging.Println("Strapi comvert is error.", err.Error())
		}
		c.JSON(http.StatusOK, strapiResponse)
	}
}

func ResetPassword(c *gin.Context) {
	requestStrapi := models.ServiceIncomeRequest{}
	jsonRequest, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Println("Request is error.", err.Error())
	}
	requestStrapi.Path = os.Getenv(constants.PATH_STRAPI_RESET_PASSWORD)
	requestStrapi.Method = constants.POST
	requestStrapi.Body = jsonRequest
	responseStrapi := services.GetInformationFormStrapi(requestStrapi)
	if responseStrapi.Status {
		c.JSON(http.StatusOK, &models.ResetPasswordResponse{
			Data: models.ResetPasswordResponseDetail{
				Status:  responseStrapi.Status,
				Message: "Your user's password has been reset.",
			},
		})
	} else {
		strapiResponse := models.StrapiError{}
		if err := json.Unmarshal([]byte(responseStrapi.Payload), &strapiResponse); err != nil {
			logging.Println("Strapi comvert is error.", err.Error())
		}
		c.JSON(http.StatusOK, strapiResponse)
	}
}

func Register(c *gin.Context) {
	requestStrapi := models.ServiceIncomeRequest{}
	jsonRequest, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Println("Request is error.", err.Error())
	}
	requestStrapi.Path = os.Getenv(constants.PATH_STRAPI_REGISTER)
	requestStrapi.Method = constants.POST
	requestStrapi.Body = jsonRequest
	responseStrapi := services.GetInformationFormStrapi(requestStrapi)
	if responseStrapi.Status {
		registerStrapiResponse := models.RegisterStrapiResponse{}
		if err := json.Unmarshal([]byte(responseStrapi.Payload), &registerStrapiResponse); err != nil {
			logging.Println("Strapi comvert is error.", err.Error())
		}
		c.JSON(http.StatusOK, registerStrapiResponse)
	} else {
		strapiResponse := models.StrapiError{}
		if err := json.Unmarshal([]byte(responseStrapi.Payload), &strapiResponse); err != nil {
			logging.Println("Strapi comvert is error.", err.Error())
		}
		c.JSON(http.StatusOK, strapiResponse)
	}
}

func CallbackProviderLogin(c *gin.Context) {
	query := c.Request.URL.Query()
	for key, val := range query {
		fmt.Println("key = " + key + ";value = " + strings.Join(val, ", "))
	}
	c.Redirect(http.StatusMovedPermanently, "https://cms.unisun.dynu.com")
}
