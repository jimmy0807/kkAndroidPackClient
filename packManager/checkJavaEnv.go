package packManager

import (
	"fmt"
	"io/ioutil"
	"kkAndroidPackClient/http/request"
	"kkAndroidPackClient/tools/zip"
	"os"
)

func ensureJavaEnv() bool {

	if !javaEnvExist() {
		request.DownloadJavaEnv()

		if !javaEnvExist() {

			return false
			//需要数据库插入日志记录了
		}
	}

	return true
}

func javaEnvExist() bool {
	dir, err := ioutil.ReadDir("./")

	if err != nil {
		return false
	}

	for _, fi := range dir {
		if !fi.IsDir() {
			if fi.Name() == "JavaEnv.zip" {
				err = zip.Unzip("JavaEnv.zip", "./")
				if err != nil {
					err := os.RemoveAll("JavaEnv")
					if err != nil {
						fmt.Println(err)
					}
					err = os.Remove("JavaEnv.zip")
					if err != nil {
						fmt.Println(err)
					}

					return false
				}

				err = os.Remove("JavaEnv.zip")
				if err != nil {
					fmt.Println(err)
				}

				return true
			}
		}
	}

	for _, fi := range dir {
		if fi.IsDir() {
			if fi.Name() == "JavaEnv" {
				return true
			}
		}
	}

	return false
}
