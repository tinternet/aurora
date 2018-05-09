package aurora

// Processor interface wraps market data processing functions.
type Processor interface {
	// Init method is executed after symbols are synchronized and
	// before other market data is downloaded and event subscriptions are done.
	// Returning error causes panic.
	Init(Context) error

	// Start method is executed after market data is synchronized and event subscriptions are done.
	// Returning error causes panic.
	Start(Context) error

	// ProcessSymbolsUpdate is executed when market symbols are updated.
	ProcessSymbolsUpdate([]Symbol)

	// ProcessOrderBook is executed when order book update is available.
	ProcessOrderBook(MarketID, OrderBook)

	// ProcessTicker is executed when ticker update is available.
	ProcessTicker(MarketID, Ticker)

	// Flush is executed when application closes.
	// It is used to finish data processing, save work if needed and gracefully exit.
	Flush()
}
