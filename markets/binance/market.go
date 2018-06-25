package binance

import (
	"github.com/bloc4ain/aurora"
	binance "github.com/bloc4ain/go-binance"
)

// Market implements official Binance API
type Market struct {
}

// ID returns market's official name
func (m Market) ID() aurora.MarketID {
	return "Binance"
}

// Symbols returns currently traded symbols on the market
func (m Market) Symbol() ([]aurora.Symbol, error) {
	var info, err = binance.GetExchangeInfo()

	if err != nil {
		return nil, err
	}

	var res = make([]aurora.Symbol, 0)

	for _, s := range info.Symbols {
		res = append(res, aurora.Symbol{
			BaseAsset:  s.BaseAsset,
			QuoteAsset: s.QuoteAsset,
		})
	}

	return res, nil
}
