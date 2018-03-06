package request

import (
	"fmt"
	"io"
	"kkAndroidPackClient/config"
	"net/http"
	"os"
	"runtime"
)

func DownloadJavaEnv() {
	fmt.Println("begin downlaod")
	fileName := ""
	url := ""
	if runtime.GOOS == "windows" {
		url = config.ServerHost + "files/JavaEnvWindows.zip"
		fileName = "JavaEnvWindows.zip"
	} else {
		url = config.ServerHost + "files/JavaEnv.zip"
		fileName = "JavaEnv.zip"
	}

	fmt.Println("download url: " + url)

	res, err := http.Get(url)
	if err != nil {
		return
	}

	f, err := os.Create(fileName)
	if err != nil {
		return
	}

	_, err32 := io.Copy(f, res.Body)
	if err32 != nil {
		return
	}

	fmt.Println("download finished")
}
