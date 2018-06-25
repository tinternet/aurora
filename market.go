package aurora

// MarketID is unique market identifier
type MarketID string

// SymbolReader interface
type SymbolReader interface {
	Read(*[]Symbol) error
	Close() error
}

// BookReader interface
type BookReader interface {
	Read(*OrderBook) error
	Close() error
}

// TickerReader interface
type TickerReader interface {
	Read(*Ticker) error
	Close() error
}

// TradesReader interface
type TradesReader interface {
	Read(*Ticker) error
	Close() error
}

// Market wraps trading market api implementation
type Market interface {
	ID() MarketID
	SymbolReader() SymbolReader
	BookReader(Symbol) BookReader
	TickerReader(Symbol) TickerReader
	// TradesReader(Symbol) TradesReader
}
