package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func DisplayCLI(stockCode string) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	stockQuote := GetQuote(stockCode)
	if len(stockQuote) == 0 {
		fmt.Print("\nInvalid stock code entered.\n")
		os.Exit(5)
	}
	stockPeers := GetPeers(stockCode)
	stockGrade := GetGrade(stockCode)
	fmt.Printf("\n\n----------------------------\n\n")
	if stockQuote[0].ChangesPercentage > 0 {
		fmt.Printf("Name: %v (%v)\nCurrent Price: %v	Previous close: %v		Precent Change: %v%%\nYear High: %v / Year Low: %v\n\n\n", stockQuote[0].Name, stockQuote[0].Symbol, stockQuote[0].Price, stockQuote[0].PreviousClose, green(stockQuote[0].ChangesPercentage), stockQuote[0].YearHigh, stockQuote[0].YearLow)
	} else {
		fmt.Printf("Name: %v (%v)\nCurrent Price: %v	Previous close: %v		Precent Change: %v%%\nYear High: %v / Year Low: %v\n\n\n", stockQuote[0].Name, stockQuote[0].Symbol, stockQuote[0].Price, stockQuote[0].PreviousClose, red(stockQuote[0].ChangesPercentage), stockQuote[0].YearHigh, stockQuote[0].YearLow)
	}
	fmt.Printf("PE Ratio: %v\n", stockQuote[0].Pe)
	fmt.Printf("Peers: ")
	for _, s := range stockPeers[0].Peers {
		peerQuote := GetQuote(s)
		if peerQuote[0].Pe < stockQuote[0].Pe {
			fmt.Printf("%v (%v)", s, red(peerQuote[0].Pe))
		} else {
			fmt.Printf("%v (%v)", s, green(peerQuote[0].Pe))
		}
		fmt.Printf(" / ")
	}
	fmt.Printf("\n\n\n")
	fmt.Printf("What the experts say:\n")
	if len(stockGrade) == 0 {
		fmt.Printf("No expect analysis is available for this code.")
	}
	for i, _ := range stockGrade {
		fmt.Printf("	%v. %v (%v)\n		Current Rating: %v | Previous Rating: %v\n", i+1, stockGrade[i].GradingCompany, stockGrade[i].Date, stockGrade[i].NewGrade, stockGrade[i].PreviousGrade)
	}
	fmt.Printf("\n\n----------------------------\n\n")
}