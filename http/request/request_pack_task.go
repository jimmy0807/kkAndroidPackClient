package request

import (
	"fmt"
	"io/ioutil"
	"kkAndroidPackClient/config"
	"kkAndroidPackClient/db/bean"
	"net/http"
	"os"
)

type PackageAppJSONResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	App     []bean.PackageApp `json:"data,omitempty"`
}

func RequestPackTask() *PackageAppJSONResponse {
	host, err := os.Hostname()
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%s", host)
	}

	resp, err := http.Get(config.ServerHost + "fetchPackTask?hostName=" + host)
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
	//fmt.Println(jsonResponse.App[0].ApkName)
	return jsonResponse
}
