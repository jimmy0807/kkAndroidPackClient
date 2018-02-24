package request

import (
	"fmt"
	"io/ioutil"
	"kkAndroidPackClient/config"
	"kkAndroidPackClient/db/bean"
	"net/http"
)

type PackageAppJSONResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	App     []bean.PackageApp `json:"data,omitempty"`
}

func RequestPackTask() *PackageAppJSONResponse {
	resp, err := http.Get(config.ServerHost + "fetchPackTask")
	if err != nil {
		return nil
	}

	jsonResponse := new(PackageAppJSONResponse)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	decodeJSONResponse(body, jsonResponse)
	fmt.Println(jsonResponse.App[0].ApkName)
	return jsonResponse
}
