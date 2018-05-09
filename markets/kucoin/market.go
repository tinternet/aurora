package kucoin

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/bloc4ain/aurora"
)

// Market implements official Binance API
type Market struct {
}

// ID returns market's official name
func (m Market) ID() aurora.MarketID {
	return "KuCoin"
}

// Symbols returns currently traded symbols on the market
func (m Market) Symbols() ([]aurora.Symbol, error) {
	var baseAssets []string
	var quoteAssets []string
	var wg sync.WaitGroup
	var err1 error
	var err2 error

	wg.Add(2)

	go func() {
		defer wg.Done()
		var res, err = http.Get("https://api.kucoin.com/v1/open/markets")
		if err != nil {
			err1 = err
			return
		}
		defer res.Body.Close()

		var markets struct {
			Assets []string `json:"data"`
		}
		if err := json.NewDecoder(res.Body).Decode(&markets); err != nil {
			err1 = err
			return
		}
		baseAssets = markets.Assets
	}()

	go func() {
		defer wg.Done()
		var res, err = http.Get("https://api.kucoin.com/v1/market/open/coins")
		if err != nil {
			err2 = err
			return
		}
		defer res.Body.Close()

		var coins struct {
			Data []struct {
				Asset string `json:"coin"`
			} `json:"data"`
		}
		if err := json.NewDecoder(res.Body).Decode(&coins); err != nil {
			err2 = err
			return
		}
		quoteAssets = make([]string, 0)
		for _, a := range coins.Data {
			quoteAssets = append(quoteAssets, a.Asset)
		}
	}()

	wg.Wait()

	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}

	var symbols = make([]aurora.Symbol, 0)
	for _, ba := range baseAssets {
		for _, qa := range quoteAssets {
			symbols = append(symbols, aurora.Symbol{BaseAsset: ba, QuoteAsset: qa})
		}
	}

	return symbols, nil
}
