package coinsync

import (
	"time"

	"github.com/bloc4ain/aurora"
)

type marketSync struct {
	market   aurora.Market
	interval time.Duration
}

var markets = make([]aurora.Market, 0)

func AddMarket(m aurora.Market) {
	markets = append(markets, m)
}

func SyncSymbols(d time.Duration) {
	ticker := time.NewTicker(d)
	for range ticker.C {

	}
}

func Symbols() []aurora.Symbol {
	symbols := make([]aurora.Symbol, 0)
	visited := make(map[aurora.SymbolID]bool)

	for _, m := range markets {
		sarr, _ := m.Symbols()
		for _, s := range sarr {
			if visited[s.ID()] {
				continue
			}
			symbols = append(symbols, s)
			visited[s.ID()] = true
		}
	}

	return symbols
}
