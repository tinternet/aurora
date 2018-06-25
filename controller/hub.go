package controller

import (
	"errors"

	"github.com/bloc4ain/aurora"
)

// EventType represents event type as integer
type EventType int

// Event represents market event
type Event struct {
	Type   EventType
	Market aurora.MarketID
	Symbol *aurora.Symbol
	Ticker *aurora.Ticker
	Book   *aurora.OrderBook
}

// EventHandler represents event handler function
type EventHandler func(*Event)

// Subscription struct
type Subscription struct {
	hub       *hub
	handler   EventHandler
	eventType EventType
}

// Close closes the subscription
func (s *Subscription) Close() {
	s.hub.unsubscribe <- s
}

type hub struct {
	// Keep subscribed handlers
	subscriptions map[EventType]map[*Subscription]bool
	// Buffered channel used for subscription
	subscribe chan *Subscription
	// Channel used for unsubscription
	unsubscribe chan *Subscription
	// Channel used for broadcasting
	broadcast chan *Event
	// Keep last events for new subscribers
	lastEvent map[EventType]*Event
}

func newHub() *hub {
	return &hub{
		subscriptions: make(map[EventType]map[*Subscription]bool),
		subscribe:     make(chan *Subscription, 1000),
		unsubscribe:   make(chan *Subscription),
		broadcast:     make(chan *Event),
		lastEvent:     make(map[EventType]*Event),
	}
}

// Subscribe subscribes the handler for given event
func (h *hub) Subscribe(et EventType, eh EventHandler) *Subscription {
	if eh == nil {
		panic(errors.New("Cannot add nil handler"))
	}
	s := &Subscription{h, eh, et}
	h.subscribe <- s
	return s
}

// Run func
func (h *hub) Run() {
	for {
		select {
		case s := <-h.subscribe:
			if h.subscriptions[s.eventType] == nil {
				h.subscriptions[s.eventType] = make(map[*Subscription]bool)
			}
			if h.lastEvent[s.eventType] != nil {
				s.handler(h.lastEvent[s.eventType])
			}
			h.subscriptions[s.eventType][s] = true

		case s := <-h.unsubscribe:
			delete(h.subscriptions[s.eventType], s)

		case e := <-h.broadcast:
			if h.subscriptions[e.Type] == nil {
				continue
			}
			for s := range h.subscriptions[e.Type] {
				s.handler(e)
			}
			h.lastEvent[e.Type] = e
		}
	}
}
