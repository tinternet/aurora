package controller

import (
	"github.com/bloc4ain/aurora"
)

type context struct {
	markets   []aurora.Market
	processor aurora.Processor
	events    chan aurora.MarketEvent
}

func (c *context) Markets() []aurora.MarketID {
	return nil
}

func (c *context) Symbols() aurora.SymbolSnapshot {

}

func (c *context) SyncTicker(sa ...aurora.Symbol) {
	for _, m := range c.markets {
		for _, s := range sa {
			go c.readTickers(s, m)
		}
	}
}

func (c *context) SyncBook(sa ...aurora.Symbol) {
	for _, m := range c.markets {
		for _, s := range sa {
			go c.readBooks(s, m)
		}
	}
}

func (c *context) run() {
	for e := range c.events {
		c.processor.Process(e, c)
	}
}

func (c *context) readTickers(s aurora.Symbol, m aurora.Market) {
	for {
		reader := m.TickerReader(s)
		for {
			ticker := new(aurora.Ticker)
			err := reader.Read(ticker)
			c.events <- aurora.MarketEvent{
				Market: m.ID(),
				Type:   aurora.TickerUpdatedEvent,
				Ticker: ticker,
			}
		}
	}
}

func (c *context) readBooks(s aurora.Symbol, m aurora.Market) {
	for {
		reader := m.BookReader(s)
		for {
			book := new(aurora.OrderBook)
			err := reader.Read(book)
			c.events <- aurora.MarketEvent{
				Market: m.ID(),
				Type:   aurora.TickerUpdatedEvent,
				Book:   book,
			}
		}
	}
}
