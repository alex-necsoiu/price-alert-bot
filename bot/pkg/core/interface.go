package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func main() {
	filter := []Filter{}

	for {
		fmt.Println("=======================================================================================")
		fmt.Printf("### In order to fix a Price Oxcilation we need the following fields:\n")
		fmt.Printf("=======================================================================================\n\n")

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Printf(`### Insert a Currency Pair! {Example: 'ETH-USD', 'BTC-USD', 'ADA-USD'}: `)
		scanner.Scan()

		fmt.Printf("\n")
		currencyPair := strings.ToUpper(scanner.Text())

		fmt.Printf(`### Insert a Fetch Interval in Seconds! {Example: 5, 10, 30}: `)
		scanner.Scan()
		fmt.Printf("\n")

		fetchInterval, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf(`### Invalid Input! Insert a Fetch Interval in Seconds! {Example: 5, 10, 30}: `)
			scanner.Scan()
			fmt.Printf("\n")
			fetchInterval, err = strconv.Atoi(scanner.Text())
			log.Fatal(err)
		}

		fmt.Printf(`### Insert a Price Osciliation Interval! {Example: 0.005, 0.01, 0.015}: `)
		scanner.Scan()
		fmt.Printf("\n")
		priceOscilationInterval, err := decimal.NewFromString(scanner.Text())
		if err != nil {
			fmt.Printf(`### Invalid Input! Insert a Price Osciliation Interval! {Example: 0.005, 0.01, 0.015}: `)
			scanner.Scan()
			priceOscilationInterval, err = decimal.NewFromString(scanner.Text())
			log.Fatal(err)
		}

		pair := Filter{
			CurrencyPair:             currencyPair,
			FetchInterval:            fetchInterval,
			PriceOsciliationInterval: priceOscilationInterval,
		}
		filter = append(filter, pair)

		fmt.Printf("### ¿Do you want to enter on other Currency Pair? {Example: 'yes'/'no'}: ")
		scanner.Scan()
		fmt.Printf("\n\n")
		if strings.ToUpper(scanner.Text()) == "NO" {
			break
		}
		if strings.ToUpper(scanner.Text()) != "NO" || strings.ToUpper(scanner.Text()) != "YES" {
			fmt.Printf("### Invalid Input ¿Do you want to enter on other Currency Pair? Example: yes/no: ")
			scanner.Scan()
			fmt.Printf("\n\n")
		}

	}

	err := MultiplePairTicker(filter)
	if err != nil {
		log.Fatal("Error")
	}
	fmt.Println("All completed, exiting")
}
