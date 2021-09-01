package ds

import (
	"github.com/gussev/d1/core/vicin"
	"github.com/gussev/d1/object"
)

// Data_Store function used to obtain an object from the store
type Data_Store interface{
	// Get give out an object from the data store according to the title && style
	Get(title string,style string,v vicin.Vicinity) (object.Object,error)
}
