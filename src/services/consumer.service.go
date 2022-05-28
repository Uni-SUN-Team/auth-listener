package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"unisun/api/auth-listener/src/constants"
	"unisun/api/auth-listener/src/logging"
	"unisun/api/auth-listener/src/models"
	"unisun/api/auth-listener/src/utils"
)

func GetInformationFormStrapi(payloadRequest models.ServiceIncomeRequest) models.ServiceIncomeResponse {
	var serviceIncomeResponse = models.ServiceIncomeResponse{}
	url := os.Getenv(constants.HOST_STRAPI_SERVICE) + os.Getenv(constants.PATH_STRAPI_INFORMATION_GATEWAY)
	payload, err := json.Marshal(payloadRequest)
	if err != nil {
		logging.Println("Change json to byte.", err.Error())
		serviceIncomeResponse.Error = err.Error()
		return serviceIncomeResponse
	} else {
		err = nil
	}
	response := utils.HTTPRequest(url, constants.POST, payload)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logging.Println("Read response from request error.", err.Error())
		serviceIncomeResponse.Error = err.Error()
		return serviceIncomeResponse
	} else {
		err = nil
		defer response.Body.Close()
	}
	err = json.Unmarshal([]byte(body), &serviceIncomeResponse)
	if err != nil {
		logging.Println("Change byte to json response.", err.Error())
		serviceIncomeResponse.Error = err.Error()
		return serviceIncomeResponse
	} else {
		err = nil

	}
	if serviceIncomeResponse.Error != "" {
		logging.Println("", serviceIncomeResponse.Error)
	}
	return serviceIncomeResponse
}

func ProcessSignin(bodyRequest models.ServiceIncomeRequest) (models.Signin, models.SignWarning, error) {
	data := GetInformationFormStrapi(bodyRequest)
	payloadSignin := models.Signin{}
	err := json.Unmarshal([]byte(data.Payload), &payloadSignin)
	if err != nil {
		logging.Println("Change byte to json", err.Error())
		return payloadSignin, models.SignWarning{
			Error: models.SignWarningError{
				Status:  http.StatusBadRequest,
				Name:    "AppError",
				Message: data.Error,
				Error:   err,
			},
		}, err
	}
	if payloadSignin.Jwt != "" && payloadSignin.User.Confirmed && payloadSignin.User.Email != "" {
		signinCallRequest := models.SigninCallRequest{}
		if userPermission, err := GetUserPermission(payloadSignin.User.Id); err != nil {
			logging.Println("Call get user-auth-permission is Error.", err.Error())
		} else {
			if userPermission.UserId != 0 && userPermission.TokenVersion != 0 {
				if jwt, err := GenerateRefreshJWT(userPermission.TokenVersion+1, userPermission.UserId); err != nil {
					logging.Println("Generate refresh token Error.", err.Error())
					return payloadSignin, models.SignWarning{Error: models.SignWarningError{
						Error:   err,
						Name:    "AppError",
						Message: err.Error(),
						Status:  http.StatusBadRequest,
					}}, err
				} else {
					payloadSignin.JwtRefresh = jwt
					signinCallRequest.Token = jwt
					signinCallRequest.UserId = userPermission.UserId
				}
			} else {
				if jwt, err := GenerateRefreshJWT(1, payloadSignin.User.Id); err != nil {
					logging.Println("Generate refresh token Error.", err.Error())
					return payloadSignin, models.SignWarning{
						Error: models.SignWarningError{
							Status:  http.StatusBadRequest,
							Name:    "AppError",
							Message: err.Error(),
							Error:   err,
						},
					}, err
				} else {
					payloadSignin.JwtRefresh = jwt
					signinCallRequest.Token = jwt
					signinCallRequest.UserId = payloadSignin.User.Id
				}
			}
		}
		if responseCallSignin, err := CallSignIn(signinCallRequest); err != nil {
			logging.Println("Call signin is Error.", err.Error())
			return payloadSignin, models.SignWarning{
				Error: models.SignWarningError{
					Status:  http.StatusBadRequest,
					Name:    "AppError",
					Message: err.Error(),
					Error:   err,
				},
			}, err
		} else {
			if responseCallSignin.Result["confirm"] == "true" {
				return payloadSignin, models.SignWarning{}, nil
			} else {
				return payloadSignin,
					models.SignWarning{Error: models.SignWarningError{
						Error:   nil,
						Name:    "CallError",
						Message: "confirm=" + responseCallSignin.Result["confirm"] + responseCallSignin.Error,
						Status:  http.StatusBadRequest,
					}}, err
			}
		}
	} else {
		signWarning := models.SignWarning{}
		err = json.Unmarshal([]byte(data.Payload), &signWarning)
		if err != nil {
			logging.Println("Change byte to json", err.Error())
			return payloadSignin,
				models.SignWarning{Error: models.SignWarningError{
					Error:   err,
					Name:    "AppError",
					Message: err.Error(),
					Status:  http.StatusUnprocessableEntity,
				}}, err
		}
		return payloadSignin, signWarning, err
	}
}
