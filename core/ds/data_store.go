package ds

import (
	"github.com/gussev/d1/core/vicin"
	"github.com/gussev/d1/object"
)

// Data_Store function used to obtain an object from the store
type Data_Store interface{
	Get(title string,style string,v vicin.Vicinity) (object.Object,error)
}
