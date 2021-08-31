package context

import "github.com/gussev/d1/core/carrier"

// Context goroutine dependant information
type Context interface{
	CreateCarrier(data_as string)
	Carrier() carrier.Carrier
}

