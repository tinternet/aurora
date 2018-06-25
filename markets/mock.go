package markets

import (
	"encoding/json"
	"io/ioutil"
	"time"

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

type symbolReader struct {
	json   []byte
	ticker *time.Ticker
}

func (r symbolReader) Read(list *[]aurora.Symbol) error {
	<-r.ticker.C
	return json.Unmarshal(binanceSymbolsJSON, list)
}

func (r symbolReader) Close() error {
	r.ticker.Stop()
	return nil
}

// ID returns market's official name
func (m BinanceMock) ID() aurora.MarketID {
	return "Binance"
}

// SymbolReader returns currently traded symbols on the market
func (m BinanceMock) SymbolReader() aurora.SymbolReader {
	return symbolReader{
		ticker: time.NewTicker(time.Second * 1),
		json:   binanceSymbolsJSON,
	}
}

// KuCoinMock implements mocked KuCoin API
type KuCoinMock struct {
}

// ID returns market's official name
func (m KuCoinMock) ID() aurora.MarketID {
	return "KuCoin"
}

// SymbolReader returns currently traded symbols on the market
func (m KuCoinMock) SymbolReader() aurora.SymbolReader {
	return symbolReader{
		ticker: time.NewTicker(time.Second * 1),
		json:   kucoinSymbolsJSON,
	}
}
