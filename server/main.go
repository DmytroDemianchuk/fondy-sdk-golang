package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type APIResponse struct {
	Response interface{} `json:"response"`
}

func main() {
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(body))

		apiResp := APIResponse{}
		err = json.Unmarshal(body, &apiResp)
		if err != nil {
			panic(err)
		}

		fmt.Println(apiResp.Response)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
