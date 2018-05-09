package aurora

// OrderBookSnapshot represents snapshot map of order books grouped by symbol and market
type OrderBookSnapshot map[MarketID]map[SymbolID]OrderBook

// TickerSnapshot represents snapshot map of tickers grouped by symbol and market
type TickerSnapshot map[MarketID]map[SymbolID]Ticker

// Controller wraps market control system functions.
type Controller interface {
	// Market returns MarketID for given string or empty if not found
	Market(string) MarketID

	// Market returns list of all registered markets.
	Markets() []MarketID

	// Symbols returns latest symbol snapshot for given markets.
	// If no markets are provided then distinct list of symbols from all markets is returned.
	// It always returns the latest snapshot and is available in all processor methods.
	Symbols(...MarketID) []Symbol

	// OrderBook returns latest order book snapshots for all subscriptions.
	// Result is always empty if executed in processor.Init().
	OrderBook() OrderBookSnapshot

	// Ticker returns latest ticker snapshots for all subscriptions.
	// Result is always empty if executed in processor.Init().
	Ticker() TickerSnapshot

	// Panic causes system shutdown due to fatal error.
	// All processors are being .Flush() before exit.
	Panic(err error)

	// SyncOrderBook subscribes for order book updates for given market and symbols
	SyncOrderBook(MarketID, ...Symbol) error
}
