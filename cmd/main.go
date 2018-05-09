package main

import (
	"log"
	"time"

	"github.com/bloc4ain/aurora/app"
	"github.com/bloc4ain/aurora/markets"
	"github.com/bloc4ain/aurora/processing"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {
	app.AddMarket(markets.BinanceMock{})
	app.AddMarket(markets.KuCoinMock{})

	app.AddProcessor(&processing.DataKeeper{})

	app.Run(app.Config{
		SymbolSyncInterval: time.Hour * 24,
	})
}
