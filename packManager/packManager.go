package packManager

import (
	"fmt"
	"kkAndroidPackClient/http/request"
	"os/exec"
	"strconv"
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

	s := "jarsigner -digestalg SHA1 -sigalg MD5withRSA -keystore kkcredit.jks -storepass weixin_kkcredit -signedjar app-base-release_340_9_Leshi_" + strconv.Itoa(i) + ".apk test.zip appKkcredit"
	fmt.Println(s)
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

		response := request.RequestPackTask()
		if response != nil {
			ensureApkIsValid(response.App[0])
		}

		//zip.CompressZip()
		//packagexxx(90)

		// result := []string{}

		// dir, err := ioutil.ReadDir("./app-aiqiyim-release")

		// if err != nil {
		// 	return
		// }

		// for _, fi := range dir {
		// 	result = append(result, "./app-aiqiyim-release/"+fi.Name())
		// }

		// fmt.Println(result)
		// archiver.Zip.Make("output.zip", result)

		//archiver.Zip.Make("output.zip", []string{"./app-aiqiyim-release"})

		// f1, err := os.Open("./app-aiqiyim-release")
		// if err != nil {
		// }

		// var files = []*os.File{f1}
		// dest := "test.zip"
		// err = zip.Zip(files, dest)

		//zip.Unzip("app-aiqiyim-release.apk", "./")

		// fmt.Println("结束")
		//go startTimer()

		//ensureJavaEnv()
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
		//modifyTime := "1912 - 1920"
		//ensureApkIsValid(modifyTime)
	}
}
