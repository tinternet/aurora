package processing

import (
	"github.com/bloc4ain/aurora"
)

type DataKeeper struct {
	ctrl aurora.Context
}

// Init method is called before markets are synched and events are subscribed
// Returning error causes panic
func (dk *DataKeeper) Init(c aurora.Context) error {
	dk.ctrl = c
	for _, m := range c.Markets() {
		if err := c.SyncOrderBook(m, c.Symbols()...); err != nil {
			return err
		}
	}
	return nil
}

// Start method is called when all data is fetched from all markets and all events are subscribed
// Returning error causes panic
func (dk *DataKeeper) Start() error {
	return nil
}

func (dk *DataKeeper) ProcessSymbolsUpdate(s []aurora.Symbol) {

}

// ProcessOrderBook is called when order book update is available
func (dk *DataKeeper) ProcessOrderBook(m aurora.MarketID, ob aurora.OrderBook) {

}

// ProcessTicker is called when ticker update is available
func (dk *DataKeeper) ProcessTicker(m aurora.MarketID, t aurora.Ticker) {

}

func (dk *DataKeeper) Flush() {

}
