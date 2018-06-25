package aurora

// EventType represents market event type as integer
type EventType int

// MarketEvent represents market event data
type MarketEvent struct {
	Type   EventType
	Market MarketID
	Symbol *Symbol
	Ticker *Ticker
	Book   *OrderBook
}

// Event types
const (
	SymbolCreatedEvent EventType = iota
	SymbolDeletedEvent
	TickerUpdatedEvent
	BookUpdatedEvent
)
