package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/shopspring/decimal"
)

var url = "https://api.uphold.com/v0/ticker/"
var icon = "github.com/uphold/alex-necsoiu/files/alert.jpeg"

// Calla Upload API each 5 sec and alerts in case of price change
func Ticker(pair *Filter) {
	ticker := time.NewTicker(time.Second * time.Duration(pair.FetchInterval))
	var priceOscilation PriceOscillliation
	priceOscilation.FirstTime = true

	defer ticker.Stop()
	for range ticker.C {
		res, err := GetData(pair.CurrencyPair)
		if err != nil {
			log.Fatal("Error:", err)
		}
		// fmt.Printf("### 	 CURRENCYPAIR:%s\n### 		  Ask:%s\n### 		  Bid:%s\n", pair.CurrencyPair, res.Ask, res.Bid)

		finish, err := CheckPriceOscillation(*pair, *res, &priceOscilation)
		if err != nil {
			log.Fatal("Error:", err)
		}
		if finish {
			break
		}
	}
}

// Calculate Price Change in percentage %
func AlertPriceChange(newInput *PriceOscillliation, oldInput *PriceOscillliation, filter *Filter) (bool, error) {

	percentAsk := newInput.Ask.Div(oldInput.Ask).Sub(decimal.NewFromFloat(1)).RoundBank(5)
	// fmt.Printf("### Percentage Ask:%+v\n", percentAsk)

	percentBid := newInput.Bid.Div(oldInput.Bid).Sub(decimal.NewFromFloat(1)).RoundBank(5)
	// fmt.Printf("### Percentage Bid:%+v\n", percentBid)
	fmt.Printf("### 	      Current Price --> PAIR: %+v   | Ask:%+v     | Bid:%+v |\n", *&newInput.CurrencyPair, *&newInput.Ask, *&newInput.Bid)
	fmt.Printf("### 	      Inicial Price --> PAIR: %+v   | Ask:%+v     | Bid:%+v |\n", *&oldInput.CurrencyPair, *&oldInput.Ask, *&oldInput.Bid)
	fmt.Println("---------------------------------------------------------------------------------------") // Insert in DB
	fmt.Printf("###  Percentege Oscillation --> PAIR: %+v   | Ask:%+v   | Bid:%+v     |\n", *&oldInput.CurrencyPair, percentAsk, percentBid)
	fmt.Println("---------------------------------------------------------------------------------------") // Insert in DB

	msg := fmt.Sprint("    ALERT PRICE OF ", *&newInput.CurrencyPair)

	// Alert if price percentage increases
	if percentAsk.Abs().GreaterThan(filter.PriceOsciliationInterval) {
		increased := " INCRESED "
		if percentAsk.IsNegative() {
			increased = " DECREASED "
		}
		msg = fmt.Sprint(msg, " ASK ", increased, percentAsk.RoundBank(5).Abs().String(), "%")
		fmt.Println(msg) // Insert in DB
		fmt.Println("#######################################################################################\n")
		beeep.Alert(filter.CurrencyPair, msg, icon)

		return true, nil
	}

	if percentBid.Abs().GreaterThan(filter.PriceOsciliationInterval) {

		increased := " INCRESED "
		if percentBid.IsNegative() {
			increased = " DECREASED "
		}
		msg = fmt.Sprint(msg, " BID ", increased, percentBid.RoundBank(5).Abs().String(), "%")

		fmt.Println(msg) // Insert in DB
		fmt.Println("#######################################################################################\n")
		beeep.Alert(filter.CurrencyPair, msg, icon)

		return true, nil
	}
	return false, nil
}

// Call Upload API and returns *Response and error
func GetData(pair string) (*Response, error) {
	client := &http.Client{}
	var res Response
	newUrl := fmt.Sprint(url, pair)

	req, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyText, &res)
	if err != nil {
		log.Fatal("Error:", err)
	}
	return &res, nil
}

// Check if the Price Oscilliation
func CheckPriceOscillation(filter Filter, input Response, obj *PriceOscillliation) (bool, error) {
	// fmt.Println("### PriceOscilation 1:", obj.FirstTime)

	bid, err := decimal.NewFromString(input.Bid)
	if err != nil {
		return false, err
	}
	ask, err := decimal.NewFromString(input.Ask)
	if err != nil {
		return false, err
	}
	if ask.IsZero() || ask.IsNegative() {
		return false, nil /// ERRORR ALGUNA COSA
	}
	if bid.IsZero() || bid.IsNegative() {
		return false, nil /// ERRORR ALGUNA COSA
	}

	newRow := PriceOscillliation{
		Ask:          ask,
		Bid:          bid,
		Currency:     input.Currency,
		CurrencyPair: filter.CurrencyPair,
		Timestamp:    time.Now(),
	}

	if obj.FirstTime {
		fmt.Println("### First Time:", obj.FirstTime, " Pair:", filter.CurrencyPair)
		obj.Ask = ask
		obj.Bid = bid
		obj.CurrencyPair = filter.CurrencyPair
		obj.Currency = input.Currency
		obj.Timestamp = time.Now()
		obj.FirstTime = false
		// fmt.Println("### After Update:", obj, " Pair:", filter.CurrencyPair)
		return false, nil
	}

	finish, err := AlertPriceChange(&newRow, obj, &filter)
	if err != nil {
		return false, err
	}

	// fmt.Printf("### Length: %+v\n\n", len(*obj))
	return finish, nil
}
