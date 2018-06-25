package main

import (
	"github.com/bloc4ain/aurora/controller"
	"github.com/bloc4ain/aurora/markets"
)

func main() {
	controller.AddMarket(markets.BinanceMock{})
	controller.AddMarket(markets.KuCoinMock{})

	controller.AddProcessor(coindata.NewProcessor())

	controller.Run()
}
