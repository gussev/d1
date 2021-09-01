package context

import "github.com/gussev/d1/core/carrier"

// Context goroutine dependant information
type Context interface{
	// CreateCarrier new Carrier
	CreateCarrier(data_as string)
	// Carrier give out a carrier
	Carrier() carrier.Carrier
}

