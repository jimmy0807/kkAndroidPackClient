package packManager

import (
	"fmt"
	"io/ioutil"
	"os/exec"
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

//Instance 获取对象
func Instance() *Manager {
	once.Do(func() {
		instance = &Manager{}

		go startTimer()
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
