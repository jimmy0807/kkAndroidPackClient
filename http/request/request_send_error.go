package request

import (
	"fmt"
	"kkAndroidPackClient/config"
	"net/http"
	"os"
)

func RequestSendError(desc string) {
	host, err := os.Hostname()
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%s", host)
	}

	_, err = http.Get(config.ServerHost + "sendError?hostName=" + host + "&desc=" + desc)
	if err != nil {
		return
	}
}
