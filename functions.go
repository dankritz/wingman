package main

import (
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
)

func DoAPICALL(url string) []byte {
	resp, getErr := http.Get(url)
    if getErr != nil {
      log.Fatal(getErr)
    }
	body, readErr := ioutil.ReadAll(resp.Body)
    if readErr != nil {
      log.Fatal(readErr)
    }
	return body
}

func GetQuote(code string) []Stock {
	apiEndpoint := "/quote/"
	url := BASEURL + apiEndpoint + code + "?apikey=" + APIKEY
	r := DoAPICALL(url)
    newquote := []Stock{}
    jsonErr := json.Unmarshal(r, &newquote)
    if jsonErr != nil {
       log.Fatal(jsonErr)
    }
	return newquote
}