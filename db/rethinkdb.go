package db

import (
	r "gopkg.in/gorethink/gorethink.v4"
)

// RethinkDB is rethink database connector
type RethinkDB struct {
	session *r.Session
}

// Connect establishes connection to rethink database
func (db *RethinkDB) Connect(url string) error {
	var err error
	db.session, err = r.Connect(r.ConnectOpts{Address: url})
	return err
}
