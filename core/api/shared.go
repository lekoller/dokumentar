package api

import (
	"encoding/json"
	"io"
	"log"
)

type APIError struct {
	Error string `json:"error"`
}

func ParseBody(reqBody io.ReadCloser, parsedBody interface{}) (err error) {
	var body []byte

	body, err = io.ReadAll(reqBody)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

// func getParams(requestURL *url.URL, keys ...string) (values []string) {
// 	queryMap := requestURL.Query()

// 	for i := 0; i < len(keys); i++ {
// 		values = append(values, "")
// 	}

// 	for mapKey, value := range queryMap {
// 		for index, key := range keys {
// 			if mapKey == key {
// 				values[index] = value[0]
// 			}
// 		}
// 	}
// 	return
// }
