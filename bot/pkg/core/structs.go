package core

import (
	"time"

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
