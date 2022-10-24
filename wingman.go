package main

import (
	//"log"
	//"net/http"
	"fmt"
	"os"
)

func main() {
	if !(ReadAPIKey()) {
		fmt.Printf("Error reading API key. Please ensure your financialmodelingprep.com API key is saved in the api.key file.\n")
		os.Exit(3)
	}

	stockCode := "SWMA.ST"
	stockQuote := GetQuote(stockCode)
	fmt.Printf("Name: %v\nPrice: %v\n\nPE Ratio: %v\n", stockQuote[0].Name, stockQuote[0].Price, stockQuote[0].Pe)

}