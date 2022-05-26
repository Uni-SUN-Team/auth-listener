package services

import (
	"encoding/json"
	"io/ioutil"
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
