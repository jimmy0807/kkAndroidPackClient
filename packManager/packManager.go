package packManager

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Manager struct {
	lastBuildNumber int
	status          string
}

var instance *Manager
var once sync.Once

func packagexxx(i int) {

	fmt.Println(i)

	s := "jarsigner -digestalg SHA1 -sigalg MD5withRSA -keystore kkcredit.jks -storepass weixin_kkcredit -signedjar app-base-release_340_9_Leshi_" + strconv.Itoa(i) + ".apk app-base-release_340_9_Leshi.apk appKkcredit"
	cmd := exec.Command("/bin/sh", "-c", s)
	cmd.Output()

	fmt.Println("finsih")

}

//Instance 获取对象
func Instance() *Manager {
	// i := 0
	// for {
	// 	if i < 1 {
	// 		go packagexxx(i)
	// 		i = i + 1
	// 	}

	// }

	once.Do(func() {
		instance = &Manager{}
		// f1, err := os.Open("./JavaEnv")
		// if err != nil {
		// }

		// var files = []*os.File{f1}
		// dest := "test.zip"
		// err = zip.Zip(files, dest)

		// zip.Unzip("./test.zip", ".")
		// fmt.Println("结束")
		//go startTimer()

		ensureJavaEnv()
	})
	return instance
}

func startTimer() {
	timer := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-timer.C:
			dealPackage()
		}
	}

	dealPackage()
}

func dealPackage() {
	for {
		modifyTime := "1912 - 1920"
		ensureApkIsValid(modifyTime)
	}
}

func ensureApkIsValid(modifyTime string) {
	if !checkApkIsVaild(modifyTime) {
		modifyApkModifyTime(modifyTime)
		removeLocalApkAndDictionary()
	}
}

func checkApkIsVaild(modifyTime string) bool {
	b, err := ioutil.ReadFile("pack.info")
	if err != nil {
		fmt.Print(err)
		return true
	}

	str := string(b)

	if str == modifyTime {
		return true
	}

	return false
}

func modifyApkModifyTime(modifyTime string) {
	s := []byte(modifyTime)
	ioutil.WriteFile("pack.info", s, 0644)
}

func removeLocalApkAndDictionary() {
	dir, err := ioutil.ReadDir("./")

	if err != nil {
		return
	}

	apkName := ""

	for _, fi := range dir {
		if !fi.IsDir() {
			if strings.HasSuffix(strings.ToLower(fi.Name()), "apk") {
				apkName = fi.Name()
			}
		}
	}

	rmrf := "rm -rf " + apkName
	cmd := exec.Command("/bin/sh", "-c", rmrf)
	cmd.Output()

	rmrfDir := "rm -rf " + strings.Split(apkName, ".")[0]
	cmd = exec.Command("/bin/sh", "-c", rmrfDir)
	cmd.Output()
}
