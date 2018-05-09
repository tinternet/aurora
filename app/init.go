package app

import (
	"log"
	"sync"
	"time"

	"github.com/bloc4ain/aurora"
)

func syncSymbols() {
	log.Printf("Syncing symbols for %d markets", len(markets))

	type symbolResult struct {
		Market  aurora.MarketID
		Symbols []aurora.Symbol
		Error   error
	}

	var results = make(chan symbolResult, len(markets))
	var wg sync.WaitGroup
	var st = time.Now()
	var retriesLimit = 5

	var fetch = func(m aurora.Market) {
		var result symbolResult
		for i := 0; i < retriesLimit; i++ {
			result.Symbols, result.Error = m.Symbols()
			result.Market = m.ID()
			if result.Error == nil {
				break
			}
		}
		results <- result
		wg.Done()
	}

	wg.Add(len(markets))

	go func() {
		for _, m := range markets {
			go fetch(m)
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var sl = make(map[aurora.MarketID][]aurora.Symbol)
	for r := range results {
		if r.Error == nil {
			sl[r.Market] = r.Symbols
		} else {
			log.Printf("Could not fetch symbols for market [%s]: %s", r.Market, r.Error)
		}
	}
	setSymbols(sl)

	log.Printf("Sync completed in %s", shortDuration(time.Since(st)))
}

func initServices() {
	var services = make([]service, len(processors))

	for i, p := range processors {
		var c = context{}
		var s = service{ctx: &c, proc: p}
		services[i] = s
	}

	for _, s := range services {
		s.proc.Init(s.ctx)
	}
}
