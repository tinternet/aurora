package controller

import (
	"errors"

	"github.com/bloc4ain/aurora"
)

// SubscribeMarket fetches symbol snapshot and subscribes for updates
func SubscribeMarket(m aurora.Market) error {
	if m == nil {
		return errors.New("Cannot add nil market")
	}
	ss, err := m.Symbols()
}

func HandleSymbols() {

}
