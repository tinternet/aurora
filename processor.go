package aurora

// Processor interface wraps market data processing functions.
type Processor interface {
	// Start method is executed after market data is synchronized and event subscriptions are done.
	Start(Context)

	// Process processes market event
	Process(MarketEvent, Context)

	// Stop is executed when application closes.
	// It is used to finish data processing, save work if needed and gracefully exit.
	Stop()
}
