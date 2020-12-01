package go_api

import (
	"encoding/json"
	"net/http"
)

type TestResponse struct {
	Result string `json:"result"`
	Message string `json:"message"`
}

func Test(w http.ResponseWriter, r *http.Request)  {
	_, client := InitializeHttpFunc(&w)
	defer client.Close()

	var apiResp TestResponse
	apiResp.Result = OK
	
	bytes, _ := json.Marshal(apiResp)
	_, _ = w.Write(bytes)
}