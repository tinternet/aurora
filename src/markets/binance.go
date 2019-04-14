package main

import (
	"net/url"

	"github.com/gorilla/websocket"
)

type binance struct{}

func (b *binance) Symbols() []string {
	return nil
}

func (b *binance) StreamTickers(channel chan<- *ticker) error {
	u := url.URL{Scheme: "wss", Host: "stream.binance.com:9443", Path: "/ws/!miniTicker@arr"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		return err
	}

	defer conn.Close()

	for {
		var resp []struct {
			EventType   string  `json:"e"`
			EventTime   int64   `json:"E"`
			Symbol      string  `json:"s"`
			ClosePrice  float64 `json:"c,string"`
			OpenPrice   float64 `json:"o,string"`
			HighPrice   float64 `json:"h,string"`
			LowPrice    float64 `json:"l,string"`
			BaseVolume  float64 `json:"v,string"`
			QuoteVolume float64 `json:"q,string"`
		}

		if err := conn.ReadJSON(&resp); err != nil {
			return err
		}

		for i := range resp {
			channel <- &ticker{
				EventTime:   resp[i].EventTime,
				Symbol:      resp[i].Symbol,
				Market:      "binance",
				ClosePrice:  resp[i].ClosePrice,
				OpenPrice:   resp[i].OpenPrice,
				HighPrice:   resp[i].HighPrice,
				LowPrice:    resp[i].LowPrice,
				BaseVolume:  resp[i].BaseVolume,
				QuoteVolume: resp[i].QuoteVolume,
			}
		}
	}
}
