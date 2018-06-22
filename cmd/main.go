package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bloc4ain/aurora/coinsync"
	"github.com/bloc4ain/aurora/markets"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func test(asd ...string) {

}

func main() {
	binance := markets.BinanceMock{}
	kucoin := markets.KuCoinMock{}
	coinsync.AddMarket(binance)
	coinsync.AddMarket(kucoin)

	// coinsync.SyncSymbols(time.Hour)

	// ticker := time.NewTicker(time.Second)

	ticker1 := make(chan struct{})
	stopped := false
	go func() {
		for !stopped {
			time.Sleep(time.Second)
			ticker1 <- struct{}{}
		}
		close(ticker1)
	}()

	go func() {
		time.Sleep(time.Second * 4)
		stopped = true
	}()

	for range ticker1 {
		fmt.Println("sync...")
	}

	fmt.Println("sync stop...")

	// app.AddProcessor(&processing.DataKeeper{})

	// app.Run(app.Config{
	// 	SymbolSyncInterval: time.Hour * 24,
	// })
}
