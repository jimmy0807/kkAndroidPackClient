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
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mholt/archiver"
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

		if ensureJavaEnv() {
			response := request.RequestPackTask()
			if response != nil {
				if ensureApkIsValid(response.App[0]) {

				}

				pack(response.App[0])
			}
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

}

func pack(app bean.PackageApp) {
	//modifyAndroidManifest(app)
	//moveAndroidManifest(app)
	removeMETAINF(app)
	renameChannel(app)
	doZip(app)
	doPack(app)
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
	targtZip := "app-" + app.ChannelName + "-release.zip"
	targtApk := "app-" + app.ChannelName + "-release.apk"
	s := "./JavaEnv/bin/jarsigner -digestalg SHA1 -sigalg MD5withRSA -keystore kkcredit.jks -storepass weixin_kkcredit -signedjar " + targtApk + " " + targtZip + " appKkcredit"
	sh.ExecuteShell(s)
	fmt.Println(s)
}
