package request

import (
	"io"
	"net/http"
	"os"
)

func RequestPackTask(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("JavaEnv.zip")
	if err != nil {
		panic(err)
	}

	_, err32 := io.Copy(f, res.Body)
	if err32 != nil {
		panic(err)
	}
}
