package packManager

import (
	"fmt"
	"io"
	"io/ioutil"
	"kkAndroidPackClient/db/bean"
	"kkAndroidPackClient/http/request"
	"kkAndroidPackClient/tools/sh"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/axgle/mahonia"
	"github.com/mholt/archiver"
)

type Manager struct {
	lastBuildNumber int
	status          string
}

var instance *Manager
var once sync.Once
var timer *time.Ticker

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
	once.Do(func() {
		instance = &Manager{}

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

		runloop()
	})
	return instance
}

func startTimer() {
	fmt.Println("启动定时器")
	timer = time.NewTicker(60 * time.Second)
	for {
		select {
		case <-timer.C:
			runloop()
		}
	}
}

func stopTimer() {
	if timer != nil {
		fmt.Println("stopTimer")
		timer.Stop()
		timer = nil
	}
}

func runloop() {
	if !dealPackage() {
		startTimer()
	} else {
		runloop()
	}
}

func dealPackage() bool {
	fmt.Println("开始打包")
	stopTimer()
	if ensureJavaEnv() {
		fmt.Println("Java环境已经下载")
		response := request.RequestPackTask()
		fmt.Println("打包数据已经返回")
		if response != nil {
			if len(response.App) == 1 {
				if ensureApkIsValid(response.App[0]) {
					return pack(response.App[0])
				}
			}
		}
	}

	return false
}

func pack(app bean.PackageApp) bool {
	removeMETAINF(app)
	renameChannel(app)
	doZip(app)
	doPack(app)
	if checkApkIsVaildFromShell(app) {
		filename := "app-" + app.ChannelName + "-release.apk"
		if request.PostFile(filename, app.ChannelID) == nil {
			removeApk(app)
			return true
		}
	}

	return false
}

func removeMETAINF(app bean.PackageApp) {
	rmrfDir := strings.Split(app.ApkName, ".")[0] + "/META-INF"
	err := os.RemoveAll(rmrfDir)
	if err != nil {
		fmt.Println(err)
	}
}

func renameChannel(app bean.PackageApp) {
	dir := strings.Split(app.ApkName, ".")[0] + "/assets/abc.txt"
	s := []byte(app.ChannelName)
	ioutil.WriteFile(dir, s, 0644)
}

func modifyAndroidManifest(app bean.PackageApp) {
	// 修改manifest
	buf, err := ioutil.ReadFile("AndroidManifest.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(buf)
	fmt.Println(content)

	a := content[8240:8270]
	var data []byte = []byte(a)
	fmt.Print(data)
	fmt.Println(strings.Index(content, "a\x00i\x00q\x00i\x00y\x00i\x00m"))

	newContent := strings.Replace(content, "\x00U\x00M\x00E\x00N\x00G\x00_\x00C\x00H\x00A\x00N\x00N\x00E\x00L\x00\x00\x00\x07\x00a\x00i\x00q\x00i\x00y\x00i\x00m", "\x00U\x00M\x00E\x00N\x00G\x00_\x00C\x00H\x00A\x00N\x00N\x00E\x00L\x00\x00\x00\x08\x00a\x00i\x00q\x00t\x00y\x00i\x00p\x00c", -1)
	// newContent := strings.Replace(content, "\x00U\x00M\x00E\x00N\x00G\x00_\x00C\x00H\x00A\x00N\x00N\x00E\x00L\x00\x00\x00", "\x00U\x00M\x00E\x00N\x00G\x00_\x00C\x00H\x00A\x00N\x00N\x00E\x00L\x00\x00\x00\x00b\x00a\x00w\x00t", -1)
	ioutil.WriteFile("AndroidManifest.xml", []byte(newContent), 0)
}

func moveAndroidManifest(app bean.PackageApp) {
	dir := strings.Split(app.ApkName, ".")[0]
	f, err := os.Create(dir + "/" + "AndroidManifest.xml")
	if err != nil {
		panic(err)
	}

	src, err := os.Open("AndroidManifest.xml")
	if err != nil {
		panic(err)
	}

	_, err32 := io.Copy(f, src)
	if err32 != nil {
		panic(err)
	}
}

func doZip(app bean.PackageApp) {
	apkDir := strings.Split(app.ApkName, ".")[0]
	targtZip := "app-" + app.ChannelName + "-release.zip"
	result := []string{}

	dir, err := ioutil.ReadDir("./" + apkDir)

	if err != nil {
		return
	}

	for _, fi := range dir {
		result = append(result, "./"+apkDir+"/"+fi.Name())
	}

	fmt.Println(result)
	archiver.Zip.Make(targtZip, result)
}

func doPack(app bean.PackageApp) {
	fmt.Println("dopack start")
	fileName := ""

	targtZip := "app-" + app.ChannelName + "-release.zip"
	targtApk := "app-" + app.ChannelName + "-release.apk"

	s := ""

	if runtime.GOOS == "windows" {
		fileName = "cd JavaEnvWindows/bin/ && jarsigner.exe "
		s = fileName + "-digestalg SHA1 -sigalg MD5withRSA -keystore kkcredit.jks -storepass weixin_kkcredit -signedjar " + "../../" + targtApk + " " + "../../" + targtZip + " appKkcredit"
	} else {
		fileName = "./JavaEnv/bin/jarsigner "
		s = fileName + "-digestalg SHA1 -sigalg MD5withRSA -keystore ./JavaEnv/bin/kkcredit.jks -storepass weixin_kkcredit -signedjar " + targtApk + " " + targtZip + " appKkcredit"
	}

	sh.ExecuteShell(s)
}

func checkApkIsVaildFromShell(app bean.PackageApp) bool {
	targtApk := "app-" + app.ChannelName + "-release.apk"
	s := ""
	if runtime.GOOS == "windows" {
		s = "cd JavaEnvWindows/bin/ && jarsigner.exe -verify " + "../../" + targtApk
	} else {
		s = "./JavaEnv/bin/jarsigner -verify " + targtApk
	}

	fmt.Println("将要验证APK")
	result := sh.ExecuteShellWithResultString(s)
	fmt.Println("验证结束")

	if runtime.GOOS == "windows" {
		dec := mahonia.NewDecoder("gbk")
		result = dec.ConvertString(result)
	}

	fmt.Println(result)
	if strings.Contains(result, "jar 已验证") {
		return true
	}

	return false
}

func removeApk(app bean.PackageApp) {
	targtApk := "app-" + app.ChannelName + "-release.apk"
	err := os.Remove(targtApk)
	if err != nil {
		fmt.Println(err)
	}

	targtZip := "app-" + app.ChannelName + "-release.zip"
	err = os.Remove(targtZip)
	if err != nil {
		fmt.Println(err)
	}
}
