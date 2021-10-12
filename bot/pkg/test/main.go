package test

import (
	"fmt"
	"sync"

	"github.com/shopspring/decimal"
)

func main() {

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

	var wg sync.WaitGroup
	wg.Add(len(filter))

	for _, row := range filter {
		go func(row Filter, wg *sync.WaitGroup) {
			defer wg.Done()
			Ticker(&row)
		}(row, &wg)
	}

	wg.Wait()
	fmt.Println("All completed, exiting")
}
