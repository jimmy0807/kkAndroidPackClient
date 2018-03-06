package packManager

import (
	"fmt"
	"io/ioutil"
	"kkAndroidPackClient/http/request"
	"kkAndroidPackClient/tools/zip"
	"os"
	"runtime"
)

func ensureJavaEnv() bool {

	if !javaEnvExist() {
		fmt.Println("javaEnvExist not exist")
		request.DownloadJavaEnv()

		if !javaEnvExist() {
			fmt.Println("javaEnvExist 还是没有")
			return false
			//需要数据库插入日志记录了
		}
	}

	return true
}

func javaEnvExist() bool {
	fileName := ""
	if runtime.GOOS == "windows" {
		fileName = "JavaEnvWindows"
	} else {
		fileName = "JavaEnv"
	}

	dir, err := ioutil.ReadDir("./")

	if err != nil {
		return false
	}

	for _, fi := range dir {
		if !fi.IsDir() {
			if fi.Name() == fileName+".zip" {
				err = zip.Unzip(fileName+".zip", "./")
				if err != nil {
					err := os.RemoveAll(fileName)
					if err != nil {
						fmt.Println(err)
					}
					err = os.Remove(fileName + ".zip")
					if err != nil {
						fmt.Println(err)
					}

					return false
				}

				err = os.Remove(fileName + ".zip")
				if err != nil {
					fmt.Println(err)
				}

				return true
			}
		}
	}

	for _, fi := range dir {
		if fi.IsDir() {
			if fi.Name() == fileName {
				return true
			}
		}
	}

	return false
}
