package coindata

// Symbol struct
type Symbol struct {
	ID    string  `gorethink:"id"`
	Base  string  `gorethink:"base"`
	Quote string  `gorethink:"quote"`
	Price float64 `gorethink:"price"`
}

// Market struct
type Market struct {
	ID      string   `gorethink:"id,omitempty"`
	Title   string   `gorethink:"title"`
	Symbols []Symbol `gorethink:"author_ids,reference" gorethink_ref:"id"`
}
