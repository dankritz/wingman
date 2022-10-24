package main

import (
	"io/ioutil"
)

var APIKEY string = "" // This will be imported from ./api.key by ReadAPIKey()
const BASEURL string = "https://financialmodelingprep.com/api/"

// Define a struct for the data we will get when requesting a Stock Quiot
type Stock struct {
	Symbol string 				`json:"symbol"`
	Name string 				`json:"name"`
	Price float32 				`json:"price"`
	ChangesPercentage float32 	`json:"changesPercentage"`
	Change float32 				`json:"change"`
	DayLow float32 				`json:"dayLow"`
	DayHigh float32 			`json:"dayHigh"`
	YearHigh float32 			`json:"yearHigh"`
	YearLow float32 			`json:"yearLow"`
	MarketCap float32 			`json:"marketCap"`
	PriceAvg50 float32 			`json:"priceAvg50"`
	PriceAvg200 float32 		`json:"priceAvg200"`
	Volume uint 				`json:"volume"`
	AvgVolume uint 				`json:"avgVolume"`
	Exchange string 			`json:"exchange"`
	Open float32 				`json:"open"`
	PreviousClose float32 		`json:"PreviousClose"`
	Eps float32 				`json:"eps"`
	Pe float32 					`json:"pe"`
	EarningsAnnouncement string `json:"earningsAnnoucement"`
	Timestamp uint32 			`json:"timestamp"`
}

type Peers struct {
	Symbol string 				`json:"symbol"`
	Peers []string				`json:"peersList"`
}

type Grade struct {
	Symbol string				`json:"symbol"`
	Date string					`json:"date"`
	GradingCompany string		`json:"gradingCompany"`
	PreviousGrade string		`json:"previousGrade"`
	NewGrade string				`json:"newGrade"`
}

func ReadAPIKey() bool {
	fileContent, err := ioutil.ReadFile("./api.key")
	if err != nil {
		return false
	}
	APIKEY = string(fileContent)
	return true
}