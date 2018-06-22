package aurora

// SymbolSnapshot holds snapshot of symbols represented grouped by id and market id in maps.
type SymbolSnapshot map[MarketID]map[SymbolID]Symbol

// OrderBookSnapshot holds snapshot of order books grouped by symbol id and market id in maps.
type OrderBookSnapshot map[MarketID]map[SymbolID]OrderBook

// TickerSnapshot holds snapshot of tickers grouped by symbol id and market id in maps.
type TickerSnapshot map[MarketID]map[SymbolID]Ticker

// Context wraps market control system functions.
type Context interface {
	// Market returns MarketID for given string or empty if not found.
	Market(string) MarketID

	// Market returns list of all registered markets.
	Markets() []MarketID

	// Symbols returns latest symbol snapshot for given markets.
	// If no markets are provided then distinct list of symbols from all markets is returned.
	// It always returns the latest snapshot and is available in all processor methods.
	Symbols() SymbolSnapshot

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
	// If no symbols are specified then subscribes for all available on the market
	SyncOrderBook(MarketID, ...Symbol) error
}
