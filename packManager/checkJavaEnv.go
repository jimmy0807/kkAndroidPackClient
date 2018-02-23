package packManager

import (
	"io/ioutil"
	"kkAndroidPackClient/http/request"
	"kkAndroidPackClient/tools/sh"
	"kkAndroidPackClient/tools/zip"
)

func ensureJavaEnv() {

	if !javaEnvExist() {
		request.DownloadJavaEnv("http://localhost:7878/files/JavaEnv.zip")

		if !javaEnvExist() {
			//需要数据库插入日志记录了
		}
	}
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
					sh.ExecuteShell("rm -rf JavaEnv")
					sh.ExecuteShell("rm -rf JavaEnv.zip")
					return false
				}
				sh.ExecuteShell("rm -rf JavaEnv.zip")
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
