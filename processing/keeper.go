package processing

import (
	"github.com/bloc4ain/aurora"
	"github.com/bloc4ain/aurora/db"
)

// DataKeeper processor stores all trade data into database
type DataKeeper struct {
	db.RethinkDB
	context aurora.Context
}

// Init method is called before markets are synched and events are subscribed
// Returning error causes panic
func (dk *DataKeeper) Init(c aurora.Context) error {
	if err := c.SyncOrderBook(c.Market("Binance")); err != nil {
		return err
	}
	return nil
}

// Start method is called when all data is fetched from all markets and all events are subscribed
// Returning error causes panic
func (dk *DataKeeper) Start(c aurora.Context) error {
	return nil
}

// ProcessSymbolsUpdate updates symbols in database
func (dk *DataKeeper) ProcessSymbolsUpdate(s []aurora.Symbol) {

}

// ProcessOrderBook updates order books in database
func (dk *DataKeeper) ProcessOrderBook(m aurora.MarketID, ob aurora.OrderBook) {

}

// ProcessTicker updates tickers in database
func (dk *DataKeeper) ProcessTicker(m aurora.MarketID, t aurora.Ticker) {

}

// Flush does nothing
func (dk *DataKeeper) Flush() {}

// NewDataKeeper returns new data keeper processor instance
func NewDataKeeper() aurora.Processor {
	return &DataKeeper{}
}
