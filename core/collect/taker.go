package collect

// Taker function to take the right object from the collected set
type Taker interface{
	Take(c Collector) string
}
