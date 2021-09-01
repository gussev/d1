package carrier

// Carrier set of functions to work with payload
type Carrier interface {
	// Payload returns content of the Carrier
	Payload() []byte
	// Err returns error in case
	Err ()      error
	// AsStrArray returns content as a string array
	AsStrArray() []string
	// Set values for payload && err if any
	Set(payload []byte,err error)
}
