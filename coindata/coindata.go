package coindata

import (
	"github.com/bloc4ain/aurora"
	r "gopkg.in/gorethink/gorethink.v4"
)

// CoinData processor saves coin information to database
type CoinData struct {
	session *r.Session
	symbols map[Symbol]bool
}

var mergeOpts = r.InsertOpts{}

// Start connects to database and merges symbols
func (cd *CoinData) Start(ctx aurora.Context) {
	var err error

	cd.session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "coindata",
	})
	if err != nil {
		panic(err)
	}

	t := r.Table("symbols")
	for _, s := range ctx.Symbols() {
		t = t.Insert(s)
	}

	_, err = t.RunWrite(cd.session)
	if err != nil {
		panic(err)
	}
}

func (cd *CoinData) Stop(ctx aurora.Context) {
	r.DB("coindata").Table("Symbols")
}

func (cd *CoinData) Process(e aurora.MarketEvent, ctx aurora.Context) {

}
