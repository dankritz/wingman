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
	var stockCode string
	fmt.Printf("Enter a stock code or CTRL+C to quit (eg. AAPL): ")
	fmt.Scanln(&stockCode)
	DisplayCLI(stockCode)
}