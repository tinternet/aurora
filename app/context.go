package app

import (
	"errors"
	"sync"

	"github.com/bloc4ain/aurora"
)

var marketSymbols struct {
	symbols     map[aurora.MarketID][]aurora.Symbol
	symbolsLock sync.RWMutex
}

func getSymbols() map[aurora.MarketID][]aurora.Symbol {
	marketSymbols.symbolsLock.RLock()
	defer marketSymbols.symbolsLock.RUnlock()
	return marketSymbols.symbols
}

func setSymbols(s map[aurora.MarketID][]aurora.Symbol) {
	marketSymbols.symbolsLock.Lock()
	marketSymbols.symbols = s
	marketSymbols.symbolsLock.Unlock()
}

type context struct {
	initialized bool
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

func (c *context) Symbols(markets ...aurora.MarketID) []aurora.Symbol {
	var list = make([]aurora.Symbol, 0)
	var visited = make(map[string]bool)
	var symbols = getSymbols()

	if len(markets) == 0 {
		for m := range symbols {
			markets = append(markets, m)
		}
	}

	for _, m := range markets {
		for _, s := range symbols[m] {
			if visited[s.String()] {
				continue
			}
			list = append(list, s)
			visited[s.String()] = true
		}
		list = append(list, symbols[m]...)
	}

	return list
}

func (c *context) OrderBook() aurora.OrderBookSnapshot {
	return nil
}

func (c *context) Ticker() aurora.TickerSnapshot {
	return nil
}

func (c *context) SyncOrderBook(m aurora.MarketID, s ...aurora.Symbol) error {
	if c.initialized { // foolproof check - subscriptions can be done only during init process
		return nil
	}
	if m == "" {
		return errors.New("Syncing order book aborted because of empty market id")
	}
	return nil
}

func (c *context) Panic(err error) {

}
