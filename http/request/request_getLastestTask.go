package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//RequestLastBuildStatus 获取最后一次编译状态
func RequestLastBuildStatus(targetName string) map[string]interface{} {
	url := targetName + "/lastBuild/api/json"

	var dat map[string]interface{}

	resp, err := http.Get(url)
	if err != nil {
		// handle error
		return nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return nil
	}

	//result := ""
	if resp.StatusCode == 200 {
		//result = string(body)

		if err := json.Unmarshal(body, &dat); err == nil {
			// fmt.Println(dat)
			// fmt.Println(dat["result"])
		} else {
			//fmt.Println(err)
		}
	}

	//fmt.Println(result)

	return dat
}
