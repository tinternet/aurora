package main

import (
	"log"
)

type ticker struct {
	EventTime   int64   `json:"eventTime"`
	Symbol      string  `json:"symbol"`
	Market      string  `json:"market"`
	ClosePrice  float64 `json:"closePrice"`
	OpenPrice   float64 `json:"openPrice"`
	HighPrice   float64 `json:"highPrice"`
	LowPrice    float64 `json:"lowPrice"`
	BaseVolume  float64 `json:"baseVolume"`
	QuoteVolume float64 `json:"quoteVolume"`
}

func main() {
	var tickers = make(chan *ticker)

	go func() {
		var market binance
		for {
			err := market.StreamTickers(tickers)
			log.Printf("Binance error: %s", err)
		}
	}()

	go publishTickers(tickers)
	consumeTickers()
}
