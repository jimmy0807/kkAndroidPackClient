package packManager

import (
	"fmt"
	"io"
	"io/ioutil"
	"kkAndroidPackClient/db/bean"
	"kkAndroidPackClient/http/request"
	"kkAndroidPackClient/tools/zip"
	"os"
	"strings"
)

func ensureApkIsValid(app bean.PackageApp) bool {
	if !checkApkIsVaild(app) {
		if request.DownloadApkFile(app) {
			modifyApkModifyTime(app.ApkLastUpdateTime)
			if apkExist(app) {
				return true
			}

			return false
		}
	}

	return true
}

func checkApkIsVaild(app bean.PackageApp) bool {
	f, err := os.OpenFile("pack.info", os.O_RDWR|os.O_CREATE, 0644)

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Print(err)
		return true
	}

	str := string(b)

	if str == app.ApkLastUpdateTime {
		if apkExist(app) {
			return true
		}

		return false
	}

	return false
}

func modifyApkModifyTime(modifyTime string) {
	s := []byte(modifyTime)
	ioutil.WriteFile("pack.info", s, 0644)
}

func apkExist(app bean.PackageApp) bool {
	dir, err := ioutil.ReadDir("./")

	if err != nil {
		return false
	}

	for _, fi := range dir {
		if !fi.IsDir() {
			if fi.Name() == app.ApkName {
				err = zip.Unzip(app.ApkName, "./")
				if err != nil {
					fmt.Println(err)
					removeLocalApkAndDictionary(app)
					return false
				}

				err := os.Remove(app.ApkName)
				if err != nil {
					fmt.Println(err)
				}

				copyAndroidManifest(app)
				return true
			}
		}
	}

	apkDir := strings.Split(app.ApkName, ".")[0]

	for _, fi := range dir {
		if fi.IsDir() {
			if fi.Name() == apkDir {
				return true
			}
		}
	}

	return false
}

func removeLocalApkAndDictionary(app bean.PackageApp) {
	err := os.Remove(app.ApkName)
	if err != nil {
		fmt.Println(err)
	}

	rmrfDir := strings.Split(app.ApkName, ".")[0]
	err = os.RemoveAll(rmrfDir)
	if err != nil {
		fmt.Println(err)
	}
}

func copyAndroidManifest(app bean.PackageApp) {
	f, err := os.Create("AndroidManifest.xml")
	if err != nil {
		panic(err)
	}

	dir := strings.Split(app.ApkName, ".")[0]
	src, err := os.Open(dir + "/" + "AndroidManifest.xml")
	if err != nil {
		panic(err)
	}

	_, err32 := io.Copy(f, src)
	if err32 != nil {
		panic(err)
	}
}
