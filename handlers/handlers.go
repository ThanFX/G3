package handlers

import (
	"encoding/json"
	"net/http"
)

type MetaJson struct {
	Code int    `json:"code"`
	Text string `json:"text,omitempty"`
}
type ResultListJson struct {
	Total int         `json:"total"`
	Items interface{} `json:"items"`
}
type ResponseWithListJson struct {
	Meta   MetaJson       `json:"meta"`
	Result ResultListJson `json:"result"`
}
type ResponseEntityJson struct {
	Meta   MetaJson    `json:"meta"`
	Result interface{} `json:"result,omitempty"`
}

func SendJsonResponse(w http.ResponseWriter, _ *http.Request, code int, content interface{}, itemCnt int, errorText string) {
	var jsonStr []byte
	if itemCnt > 0 {
		var response ResponseWithListJson
		response.Meta.Code = code
		response.Result.Total = itemCnt
		response.Result.Items = content
		jsonStr, _ = json.Marshal(response)
	} else {
		var response ResponseEntityJson
		response.Meta.Code = code
		response.Meta.Text = errorText
		response.Result = content
		jsonStr, _ = json.Marshal(response)
	}
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.WriteHeader(code)
	w.Write(jsonStr)
}
