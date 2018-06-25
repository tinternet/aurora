package controller

import (
	"github.com/bloc4ain/aurora"
)

type Controller struct {
	markets    []aurora.Market
	processors []aurora.Processor
}

func NewController() *Controller {
	return &Controller{
		markets:    make([]aurora.Market, 0),
		processors: make([]aurora.Processor, 0),
	}
}

func (c *Controller) AddMarket(m aurora.Market) {
	c.markets = append(c.markets, m)
}

func (c *Controller) AddProcessor(p aurora.Processor) {
	c.processors = append(c.processors, p)
}

func (c *Controller) Run() {
	var symbols = make(aurora.SymbolSnapshot)
	for _, m := range c.markets {
		s := make(chan []aurora.Symbol)
		go readSymbols(s, m)
		arr := <-s
	}

	for _, p := range c.processors {
		ctx := newContext(p)
		go ctx.Run()
	}
}

func readSymbols(s chan<- []aurora.Symbol, m aurora.Market) {
	for {
		sr := m.SymbolReader()
		for {
			arr := make([]aurora.Symbol, 0)
			sr.Read(&arr)
			s <- arr
		}
	}
}
