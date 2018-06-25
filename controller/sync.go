package controller

import "github.com/bloc4ain/aurora"

type sync struct {
	symbols []aurora.SymbolID
}

func (s sync) SyncTicker(symbol aurora.Symbol) {

}

func (s sync) SyncBook(symbol aurora.Symbol) {

}

var syncBook chan aurora.SymbolID
var syncTicker chan aurora.SymbolID

func syncBooks() {
	for symbol := range syncBook {

	}
}
