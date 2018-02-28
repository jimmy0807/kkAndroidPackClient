package request

import (
	"io"
	"kkAndroidPackClient/config"
	"net/http"
	"os"
)

func DownloadJavaEnv() {
	url := config.ServerHost + "files/JavaEnv.zip"

	res, err := http.Get(url)
	if err != nil {
		return
	}

	f, err := os.Create("JavaEnv.zip")
	if err != nil {
		return
	}

	_, err32 := io.Copy(f, res.Body)
	if err32 != nil {
		return
	}
}
