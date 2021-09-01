package collect

// Taker function to take the right object from the collected set
type Taker interface{
	// Take selects the right object according to the given criteria
	Take(c Collector) string
}
