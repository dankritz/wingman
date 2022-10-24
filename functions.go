package main

import (
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
	//"fmt"
)

func DoAPICALL(url string) []byte {
	//fmt.Println(url)
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
	url := BASEURL + "v3" + apiEndpoint + code + "?apikey=" + APIKEY
	r := DoAPICALL(url)
    newQuote := []Stock{}
    jsonErr := json.Unmarshal(r, &newQuote)
    if jsonErr != nil {
       log.Fatal(jsonErr)
    }
	return newQuote
}

func GetPeers(code string) []Peers {
	apiEndpoint := "/stock_peers?"
	url := BASEURL + "v4"+ apiEndpoint + "symbol=" + code + "&apikey=" + APIKEY
	r := DoAPICALL(url)
    newPeers := []Peers{}
    jsonErr := json.Unmarshal(r, &newPeers)
    if jsonErr != nil {
       log.Fatal(jsonErr)
    }
	return newPeers
}

func GetGrade(code string) []Grade {
	apiEndpoint := "/grade/"
	url := BASEURL + "v3" + apiEndpoint + code + "?limit=5" +"&apikey=" + APIKEY
	r := DoAPICALL(url)
    newGrade := []Grade{}
    jsonErr := json.Unmarshal(r, &newGrade)
    if jsonErr != nil {
       log.Fatal(jsonErr)
    }
	return newGrade
}
