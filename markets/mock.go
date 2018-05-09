package markets

import (
	"encoding/json"
	"io/ioutil"

	"github.com/bloc4ain/aurora"
)

const mockDataDir = "markets/mockdata"

const (
	binanceSymbolsFile = mockDataDir + "/binance-symbols.json"
	kucoinSymbolsFile  = mockDataDir + "/kucoin-symbols.json"
)

var (
	binanceSymbolsJSON []byte
	kucoinSymbolsJSON  []byte
)

func init() {
	var err error
	if binanceSymbolsJSON, err = ioutil.ReadFile(binanceSymbolsFile); err != nil {
		panic(err)
	}
	if kucoinSymbolsJSON, err = ioutil.ReadFile(kucoinSymbolsFile); err != nil {
		panic(err)
	}
}

// BinanceMock implements mocked Binance API
type BinanceMock struct {
}

// ID returns market's official name
func (m BinanceMock) ID() aurora.MarketID {
	return "Binance"
}

// Symbols returns currently traded symbols on the market
func (m BinanceMock) Symbols() ([]aurora.Symbol, error) {
	var list []aurora.Symbol
	var err = json.Unmarshal(binanceSymbolsJSON, &list)
	return list, err
}

// KuCoinMock implements mocked KuCoin API
type KuCoinMock struct {
}

// ID returns market's official name
func (m KuCoinMock) ID() aurora.MarketID {
	return "KuCoin"
}

// Symbols returns currently traded symbols on the market
func (m KuCoinMock) Symbols() ([]aurora.Symbol, error) {
	var list []aurora.Symbol
	var err = json.Unmarshal(kucoinSymbolsJSON, &list)
	return list, err
}
