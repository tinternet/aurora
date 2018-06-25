package aurora

// SymbolSnapshot holds snapshot of symbols represented grouped by id and market id in maps.
type SymbolSnapshot map[MarketID]map[SymbolID]Symbol

// OrderBookSnapshot holds snapshot of order books grouped by symbol id and market id in maps.
type OrderBookSnapshot map[MarketID]map[SymbolID]OrderBook

// TickerSnapshot holds snapshot of tickers grouped by symbol id and market id in maps.
type TickerSnapshot map[MarketID]map[SymbolID]Ticker

// Context wraps market control system functions.
type Context interface {
	// Market returns list of all registered markets.
	Markets() []MarketID

	// Symbols returns latest symbol snapshot for given markets.
	// If no markets are provided then distinct list of symbols from all markets is returned.
	// It always returns the latest snapshot and is available in all processor methods.
	Symbols() SymbolSnapshot
}
