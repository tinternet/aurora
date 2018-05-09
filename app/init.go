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

	var results = make(chan symbolResult)
	var wg sync.WaitGroup
	var st = time.Now()

	var fetch = func(m aurora.Market) {
		var result symbolResult
		var retriesLimit = 5
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

	for _, m := range markets {
		wg.Add(1)
		go fetch(m)
	}

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

	// Create services
	for i, p := range processors {
		var c = context{}
		var s = service{ctx: &c, proc: p}
		services[i] = s
	}

	// Initialize processors
	for _, s := range services {
		if err := s.proc.Init(s.ctx); err != nil {
			panic(err)
		}
		s.ctx.initialized = true
	}

	// Start processors
	for _, s := range services {
		if err := s.proc.Start(s.ctx); err != nil {
			panic(err)
		}
		s.ctx.initialized = true
	}
}
