package app

import "github.com/bloc4ain/aurora"

var markets = make(map[aurora.MarketID]aurora.Market)
var processors = make([]aurora.Processor, 0)
var services = make([]service, 0)
