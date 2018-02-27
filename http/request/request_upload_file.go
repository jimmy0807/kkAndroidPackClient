package request

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"kkAndroidPackClient/config"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

/*
func main() {
    target_url := "http://localhost/upload"
    filename := "./example.pdf"
    postFile(filename, target_url)
}
*/
func PostFile(filename string, channelID int64) error {
	targetUrl := config.ServerHost + "uploadApkFile"

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	bodyWriter.WriteField("fileName", filename)
	bodyWriter.WriteField("channelID", strconv.FormatInt(channelID, 10))

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadFile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
