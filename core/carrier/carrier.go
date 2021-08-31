package carrier

// Carrier set of functions to work with payload
type Carrier interface {
	Payload() []byte
	Err ()      error
	AsStrArray() []string
	Set(payload []byte,err error)
}
