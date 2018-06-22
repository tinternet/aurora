package app

import (
	"errors"
	"sync"

	"github.com/bloc4ain/aurora"
)

var marketSymbols struct {
	symbols     aurora.SymbolSnapshot
	symbolsLock sync.RWMutex
}

func getSymbols() aurora.SymbolSnapshot {
	marketSymbols.symbolsLock.RLock()
	defer marketSymbols.symbolsLock.RUnlock()
	return marketSymbols.symbols
}

func setSymbols(s aurora.SymbolSnapshot) {
	marketSymbols.symbolsLock.Lock()
	marketSymbols.symbols = s
	marketSymbols.symbolsLock.Unlock()
}

type context struct {
	subs map[aurora.Market][]aurora.Symbol
}

func (c *context) Market(m string) aurora.MarketID {
	if m == "" {
		return ""
	}
	if v := markets[aurora.MarketID(m)]; v == nil {
		return ""
	}
	return aurora.MarketID(m)
}

func (c *context) Markets() []aurora.MarketID {
	var l = make([]aurora.MarketID, 0)
	for m := range markets {
		l = append(l, m)
	}
	return l
}

func (c *context) Symbols(markets ...aurora.MarketID) aurora.SymbolSnapshot {
	return getSymbols()
}

func (c *context) OrderBook() aurora.OrderBookSnapshot {
	return nil
}

func (c *context) Ticker() aurora.TickerSnapshot {
	return nil
}

func (c *context) SyncOrderBook(m aurora.MarketID, s ...aurora.Symbol) error {
	if m == "" {
		return errors.New("Syncing order book aborted because of empty market id")
	}
	if markets[m] == nil {
		return errors.New("Market with name " + string(m) + "not found")
	}
	if c.subs == nil {
		c.subs = make(map[aurora.Market][]aurora.Symbol)
	}
	if c.subs[markets[m]] == nil {
		c.subs[markets[m]] = make([]aurora.Symbol, 0)
	}
	if len(s) == 0 {
		for _, m := range getSymbols() {
			for _, sy := range m {
				s = append(s, sy)
			}
		}
	}
	c.subs[markets[m]] = append(c.subs[markets[m]], s...)
	return nil
}

func (c *context) Panic(err error) {

}
