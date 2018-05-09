package app

import (
	"time"

	"github.com/bloc4ain/aurora"
)

// AddMarket registers market in the system
func AddMarket(m aurora.Market) {
	if m == nil {
		panic("Cannot add nil market")
	}
	if _, exists := markets[m.ID()]; exists {
		panic(`Market with name "` + m.ID() + `" already exists`)
	}
	markets[m.ID()] = m
}

// MarketList returns list of all registered markets in the system
func MarketList() []aurora.Market {
	var list = make([]aurora.Market, 0)
	for _, m := range markets {
		list = append(list, m)
	}
	return list
}

// AddProcessor registers data processor in the system
func AddProcessor(p aurora.Processor) {
	if p == nil {
		panic("Cannot add nil processor")
	}
	processors = append(processors, p)
}

// Run starts the controller process
func Run(c Config) {
	syncSymbols()
	initServices()

	var ss = time.NewTicker(c.SymbolSyncInterval)
	for {
		select {
		case <-ss.C:
			syncSymbols()
		}
	}
}
