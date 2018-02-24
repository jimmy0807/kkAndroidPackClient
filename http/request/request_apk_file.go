package request

import (
	"fmt"
	"io"
	"kkAndroidPackClient/config"
	"kkAndroidPackClient/db/bean"
	"net/http"
	"os"
)

func DownloadApkFile(app bean.PackageApp) bool {
	url := config.ServerHost + "files/" + app.ApkName
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return false
	}

	f, err := os.Create(app.ApkName)
	if err != nil {
		fmt.Println(err)
		return false
	}

	w, err32 := io.Copy(f, res.Body)
	if err32 != nil {
		fmt.Println(err32)
		return false
	}

	fmt.Println(w)
	return true
}
