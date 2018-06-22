package aurora

import (
	"io"
)

// SymbolStream wraps symbol read method
type SymbolStream interface {
	io.Closer
	Read() ([]Symbol, error)
}

// OrderBookStream wraps order book read method
type OrderBookStream interface {
	io.Closer
	Read() ([]OrderBook, error)
}

// TickerStream wraps ticker read method
type TickerStream interface {
	io.Closer
	Read() ([]Ticker, error)
}

// TradeStream wraps trades read method
type TradeStream interface {
	io.Closer
	Read() ([]Trade, error)
}

// MarketID is unique market identifier
type MarketID string

// Market wraps trading market api implementation
type Market interface {
	ID() MarketID
	Symbols() ([]Symbol, error)
	SymbolStream() (SymbolStream, error)
	OrderBookStream(Symbol) (OrderBookStream, error)
	TickerStream(Symbol) (TickerStream, error)
	TradeStream(Symbol) (TradeStream, error)
}
