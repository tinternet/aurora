package aurora

import (
	"fmt"
	"time"
)

type SymbolID string

// Symbol holds symbol assets.
type Symbol struct {
	BaseAsset  string
	QuoteAsset string
}

func (s Symbol) String() string {
	return fmt.Sprintf("Symbol[Base:%s;Quote:%s]", s.BaseAsset, s.QuoteAsset)
}

func (s Symbol) ID() SymbolID {
	return SymbolID(s.BaseAsset + s.QuoteAsset)
}

// Ticker represents 24 hour stats for symbol.
type Ticker struct {
	Symbol      Symbol
	Time        time.Time
	Price       float64
	LowPrice    float64
	HighPrice   float64
	BaseVolume  float64
	QuoteVolume float64
}

func (t Ticker) String() string {
	return fmt.Sprintf("Ticker[Symbol:%s;Price:%.8f;LowPrice:%.8f;HighPrice:%.8f;BaseVolume:%.8f;QuoteVolume:%.8f]",
		t.Symbol, t.Price, t.LowPrice, t.HighPrice, t.BaseVolume, t.QuoteVolume)
}

// Order represents market order.
type Order struct {
	Rate     float64
	Quantity float64
}

func (o Order) String() string {
	return fmt.Sprintf("Order[Rate:%.8f;Quantity:%.8f]", o.Rate, o.Quantity)
}

// OrderBook represents market buy and sell orders.
type OrderBook struct {
	Symbol Symbol
	Buy    []Order
	Sell   []Order
}

func (ob OrderBook) String() string {
	return fmt.Sprintf("OrderBook[Symbol:%s;Buy:%v;Sell:%v", ob.Symbol, ob.Buy, ob.Sell)
}

// Trade represents completed market trade.
type Trade struct {
	Symbol   Symbol
	Price    float64
	Quantity float64
}

func (t Trade) String() string {
	return fmt.Sprintf("Trade[Symbol:%s;Price:%.8f;Quantity:%.8f", t.Symbol, t.Price, t.Quantity)
}
