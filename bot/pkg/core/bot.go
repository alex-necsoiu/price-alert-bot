package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/shopspring/decimal"
)

type Response struct {
	Ask      string `json:"ask"`
	Bid      string `json:"bid"`
	Currency string `json:"currency"`
}

type PriceOscillliation struct {
	Ask                 decimal.Decimal `json:"ask"`
	Bid                 decimal.Decimal `json:"bid"`
	PriceOscilationAsk  decimal.Decimal `json:"price_osciliation_ask"`
	PriceOsciliationBid decimal.Decimal `json:"price_osciliation_bid"`
	Currency            string          `json:"currency"`
	CurrencyPair        string          `json:"currency_pair"`
	Timestamp           time.Time       `json:"timestamp"`
	FirstTime           bool            `json:"first_time"`
}
type Filter struct {
	CurrencyPair             string          `json:"currennncy_pair"`
	FetchInterval            int             `json:"fetch_interval"`
	PriceOsciliationInterval decimal.Decimal `json:"price_osciliation_interval"`
}

var url = "https://api.uphold.com/v0/ticker/"
var icon = "github.com/alex-necsoiu/uphold-bot/bot/files/alert.jpeg"

// Requests
func MultiplePairTicker(filter []Filter) error {
	var wg sync.WaitGroup
	wg.Add(len(filter))

	for _, row := range filter {
		go func(row Filter, wg *sync.WaitGroup) {
			defer wg.Done()
			Ticker(&row)
		}(row, &wg)
	}

	wg.Wait()
	return nil
}

// Call Upload API each FetchInterval in Seconds and alerts in case of price change
func Ticker(pair *Filter) {
	var priceChange PriceOscillliation
	priceChange.FirstTime = true
	finish := false

	// Fetch Upload Api each
	ticker := time.NewTicker(time.Second * time.Duration(pair.FetchInterval))
	defer ticker.Stop()
	for range ticker.C {
		errorIntent := 0
		res, err := GetData(pair.CurrencyPair)
		if err != nil {
			// Tolerate until 3 bat responses from REST API request
			if errorIntent < 3 {
				errorIntent++
				continue
			}
			break
		}
		finish, err = CheckPriceOscillation(*pair, *res, &priceChange)
		if err != nil {
			if errorIntent < 3 {
				errorIntent++
				continue
			}
			break
		}
		if finish {
			break
		}
	}
	return
}

// Calculate Price Change in percentage %
func AlertPriceChange(newInput *PriceOscillliation, oldInput *PriceOscillliation, filter *Filter) (bool, error) {

	percentAsk := newInput.Ask.Div(oldInput.Ask).Sub(decimal.NewFromFloat(1)).RoundBank(4)

	percentBid := newInput.Bid.Div(oldInput.Bid).Sub(decimal.NewFromFloat(1)).RoundBank(4)

	msg := fmt.Sprint("    ALERT PRICE OF ", *&newInput.CurrencyPair)

	// Alert if price percentage changes
	if percentAsk.Abs().GreaterThan(filter.PriceOsciliationInterval) {
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Printf("### 	      Current Price | PAIR: %+v   | Ask:%+v     | Bid:%+v |\n", *&newInput.CurrencyPair, *&newInput.Ask, *&newInput.Bid)
		fmt.Printf("### 	      Inicial Price | PAIR: %+v   | Ask:%+v     | Bid:%+v |\n", *&oldInput.CurrencyPair, *&oldInput.Ask, *&oldInput.Bid)
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Printf("###  Percentege Oscillation | PAIR: %+v   | Ask:%+v     | Bid:%+v   |\n", *&oldInput.CurrencyPair, percentAsk, percentBid)
		fmt.Println("---------------------------------------------------------------------------------------")

		increased := " INCRESED "
		if percentAsk.IsNegative() {
			increased = " DECREASED "
		}
		msg = fmt.Sprint(msg, " ASK ", increased, percentAsk.RoundBank(5).Abs().String(), "%")
		fmt.Println(msg)
		fmt.Println("---------------------------------------------------------------------------------------") // Insert into DB

		beeep.Alert(filter.CurrencyPair, msg, icon)

		return true, nil
	}

	if percentBid.Abs().GreaterThan(filter.PriceOsciliationInterval) {
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Printf("### 	      Current Price | PAIR: %+v   | Ask:%+v     | Bid:%+v |\n", *&newInput.CurrencyPair, *&newInput.Ask, *&newInput.Bid)
		fmt.Printf("### 	      Inicial Price | PAIR: %+v   | Ask:%+v     | Bid:%+v |\n", *&oldInput.CurrencyPair, *&oldInput.Ask, *&oldInput.Bid)
		fmt.Println("---------------------------------------------------------------------------------------")
		fmt.Printf("###  Percentege Oscillation | PAIR: %+v   | Ask:%+v     | Bid:%+v   |\n", *&oldInput.CurrencyPair, percentAsk, percentBid)
		fmt.Println("---------------------------------------------------------------------------------------")

		increased := " INCRESED "
		if percentBid.IsNegative() {
			increased = " DECREASED "
		}
		msg = fmt.Sprint(msg, " BID ", increased, percentBid.RoundBank(5).Abs().String(), "%")

		fmt.Println(msg)
		fmt.Println("---------------------------------------------------------------------------------------") // Insert into DB

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
		return nil, err
	}
	return &res, nil
}

// Check if the Price Oscilliation
func CheckPriceOscillation(filter Filter, input Response, obj *PriceOscillliation) (bool, error) {
	bid, err := decimal.NewFromString(input.Bid)
	if err != nil {
		return false, err
	}
	ask, err := decimal.NewFromString(input.Ask)
	if err != nil {
		return false, err
	}
	if ask.IsZero() || ask.IsNegative() {
		return false, errors.New("Invalid Price received from API Reques!")
	}
	if bid.IsZero() || bid.IsNegative() {
		return false, errors.New("Invalid Price received from API Reques!")
	}

	// Map Response from Api to PriceOscillliation
	newRow := PriceOscillliation{
		Ask:          ask,
		Bid:          bid,
		Currency:     input.Currency,
		CurrencyPair: filter.CurrencyPair,
		Timestamp:    time.Now(),
	}

	if obj.FirstTime {
		obj.Ask = ask
		obj.Bid = bid
		obj.CurrencyPair = filter.CurrencyPair
		obj.Currency = input.Currency
		obj.Timestamp = time.Now()
		obj.FirstTime = false
		return false, nil
	}

	// Checks if the price changes more then fixed
	finish, err := AlertPriceChange(&newRow, obj, &filter)
	if err != nil {
		return false, err
	}
	return finish, nil
}
