package main

import (
	"log"
	"testing"

	"github.com/shopspring/decimal"
)

type RequestTable struct {
	TestName string `json:"test_name"`
	Filter   []Filter
}

func TestPriceOscillationBot(t *testing.T) {
	validCases, err := GetSingleCurrencyAlert()
	assertNoErr(t, err)
	RunTestsCase(t, validCases)

	validCases, err = GetMultipleCurrencyAlert()
	assertNoErr(t, err)
	RunTestsCase(t, validCases)

	invalidCases, err := GetInvalidSingleCurrencyAlert()
	assertNoErr(t, err)
	RunTestsInvalidCase(t, invalidCases)

}

// RunTestsValidCardPayins runs a test for each obj inside []cardPayinTable
func RunTestsCase(t *testing.T, validCases []RequestTable) {
	for _, x := range validCases {
		t.Run(x.TestName, func(t *testing.T) {
			err := MultiplePairTicker(x.Filter)
			if err != nil {
				t.Errorf("Error failed with error: %v", err)
			}
		})
	}
}

func RunTestsInvalidCase(t *testing.T, invalidCases []RequestTable) {
	for _, x := range invalidCases {
		t.Run(x.TestName, func(t *testing.T) {
			// res, _ := s.CardPayinCreate(req(ctx, x.cardPayinRequest))+
			err := MultiplePairTicker(x.Filter)
			if err == nil {
				t.Errorf("Should return error and returns %v", err)
			}
		})
	}
}

// GetValidCardPayins returns a array type []cardPayinTable with valid cases
func GetMultipleCurrencyAlert() ([]RequestTable, error) {
	// var currencyPair []string
	// currencyPair = append(currencyPair, "BTC-USD", "ETH-USD", "ADA-USD")
	priceOscilationInterval1 := decimal.NewFromFloat(0.0004)
	priceOscilationInterval2 := decimal.NewFromFloat(0.0007)
	priceOscilationInterval3 := decimal.NewFromFloat(0.0010)

	pair1 := Filter{
		CurrencyPair:             "BTC-USD",
		FetchInterval:            8,
		PriceOsciliationInterval: priceOscilationInterval1,
	}

	pair2 := Filter{
		CurrencyPair:             "ETH-USD",
		FetchInterval:            5,
		PriceOsciliationInterval: priceOscilationInterval2,
	}
	pair3 := Filter{
		CurrencyPair:             "ADA-USD",
		FetchInterval:            3,
		PriceOsciliationInterval: priceOscilationInterval3,
	}
	filter := []Filter{}
	filter = append(filter, pair1, pair2, pair3)

	validCases := []RequestTable{
		{TestName: "", Filter: filter},
	}
	return validCases, nil
}

// GetValidCardPayins returns a array type []cardPayinTable with valid cases
func GetSingleCurrencyAlert() ([]RequestTable, error) {
	// var currencyPair []string
	// currencyPair = append(currencyPair, "BTC-USD", "ETH-USD", "ADA-USD")
	priceOscilationInterval1 := decimal.NewFromFloat(0.0004)

	pair1 := Filter{
		CurrencyPair:             "BTC-USD",
		FetchInterval:            8,
		PriceOsciliationInterval: priceOscilationInterval1,
	}

	filter := []Filter{}
	filter = append(filter, pair1)

	validCases := []RequestTable{
		{TestName: "Test Multiple Currency Alert", Filter: filter},
	}
	return validCases, nil
}

func GetInvalidSingleCurrencyAlert() ([]RequestTable, error) {
	// var currencyPair []string
	// currencyPair = append(currencyPair, "BTC-USD", "ETH-USD", "ADA-USD")
	priceOscilationInterval1 := decimal.NewFromFloat(0.0004)

	pair1 := Filter{
		CurrencyPair:             "USD-USD",
		FetchInterval:            8,
		PriceOsciliationInterval: priceOscilationInterval1,
	}

	filter := []Filter{}
	filter = append(filter, pair1)

	validCases := []RequestTable{
		{TestName: "Test Invalid Pair", Filter: filter},
	}
	return validCases, nil
}
func assertNoErr(t *testing.T, err error) {
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
}
