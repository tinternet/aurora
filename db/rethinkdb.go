package db

import (
	r "gopkg.in/gorethink/gorethink.v4"
)

type Currency struct {
	Name              string
	Cap               float64
	Price             float64
	Volume24H         float64
	CirculatingSupply float64
	Ticker24H         float64
	PriceGraph        []float64
}

// RethinkDB is rethink database connector
type RethinkDB struct {
	session *r.Session
}

// Connect establishes connection to rethink database
func (db *RethinkDB) Connect(url string) error {
	var err error
	// db.session, err = r.Connect(r.ConnectOpts{Address: url})
	return err
}
