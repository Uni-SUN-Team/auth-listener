package services

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"unisun/api/unisun-authen-listener/src/constants"
	"unisun/api/unisun-authen-listener/src/logging"
	"unisun/api/unisun-authen-listener/src/models"
	"unisun/api/unisun-authen-listener/src/utils"
)

func GetUserPermission(userId int) (models.UserAuthPermission, error) {
	userAuthPermission := models.UserAuthPermission{}
	url := os.Getenv(constants.AUTHEN_GATEWAY_HOST) + os.Getenv(constants.AUTHEN_GATEWAY_PATH_GET_TOKENVERSION) + strconv.Itoa(userId)
	response := utils.HTTPRequest(url, constants.GET, nil)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logging.Println("Read response from request error.", err.Error())
		return userAuthPermission, err
	} else {
		err = nil
		defer response.Body.Close()
	}
	err = json.Unmarshal([]byte(body), &userAuthPermission)
	if err != nil {
		logging.Println("Change byte to json response.", err.Error())
		return userAuthPermission, err
	} else {
		err = nil

	}
	return userAuthPermission, nil
}

func CallSignIn(payloadRequest models.SigninCallRequest) (models.CallAuthenGatewayResponse, error) {
	response := models.CallAuthenGatewayResponse{}
	url := os.Getenv(constants.AUTHEN_GATEWAY_HOST) + os.Getenv(constants.AUTHEN_GATEWAY_PATH_SIGNIN)
	payload, err := json.Marshal(payloadRequest)
	if err != nil {
		logging.Println("Change json to byte.", err.Error())
		return response, err
	} else {
		err = nil
	}
	httpResponse := utils.HTTPRequest(url, constants.POST, payload)
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		logging.Println("Read response from request error.", err.Error())
		return response, err
	} else {
		err = nil
		defer httpResponse.Body.Close()
	}
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		logging.Println("Change byte to json response.", err.Error())
		return response, err
	} else {
		err = nil

	}
	return response, nil
}

func CallRevoke(payloadRequest models.Revoke) (models.CallAuthenGatewayResponse, error) {
	response := models.CallAuthenGatewayResponse{}
	url := os.Getenv(constants.AUTHEN_GATEWAY_HOST) + os.Getenv(constants.AUTHEN_GATEWAY_PATH_CALL_REVOKE)
	payload, err := json.Marshal(payloadRequest)
	if err != nil {
		logging.Println("Change json to byte.", err.Error())
		return response, err
	} else {
		err = nil
	}
	httpResponse := utils.HTTPRequest(url, constants.POST, payload)
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		logging.Println("Read response from request error.", err.Error())
		return response, err
	} else {
		err = nil
		defer httpResponse.Body.Close()
	}
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		logging.Println("Change byte to json response.", err.Error())
		return response, err
	} else {
		err = nil

	}
	return response, nil
}

func CallRefreshToken(payloadRequest models.RefreshTokenBodyRequest) (models.RefreshTokenBodyResponse, error) {
	logging.Println("Start call refresh token.", "")
	response := models.RefreshTokenBodyResponse{}
	url := os.Getenv(constants.AUTHEN_GATEWAY_HOST) + os.Getenv(constants.AUTHEN_GATEWAY_PATH_CALL_REFRESHTOKEN)
	payload, err := json.Marshal(payloadRequest)
	if err != nil {
		logging.Println("Change json to byte.", err.Error())
		return response, err
	}
	httpResponse := utils.HTTPRequest(url, constants.POST, payload)
	if httpResponse.StatusCode == 200 {
		body, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			logging.Println("Read response from request error.", err.Error())
			return response, err
		} else {
			err = nil
			defer httpResponse.Body.Close()
		}
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			logging.Println("Change byte to json response.", err.Error())
			return response, err
		}
		logging.Println("End call refresh token.", "")
		return response, nil
	} else {
		logging.Println("End call refresh token.", "Is status "+strconv.Itoa(httpResponse.StatusCode))
		return response, nil
	}
}
