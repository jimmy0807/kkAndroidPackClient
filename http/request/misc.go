package request

import (
	"encoding/json"
)

// JSONResponse 返回数据

func decodeJSONResponse(body []byte, data interface{}) {
	if err := json.Unmarshal(body, data); err == nil {
	} else {
	}
}
