package reader

import (
	"github.com/gussev/d1/core/context"
	"github.com/gussev/d1/core/vicin"
)
// Reader function to read object
type Reader interface{
	/*
		title - name of the object
		style - subtype of the object
		vicinity - environment data
	    context  - goroutine dependant data
	*/
	Read(title string,style string,v vicin.Vicinity, context context.Context)
}
