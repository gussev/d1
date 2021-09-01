package build

import (
	"github.com/gussev/d1/core/context"
	"github.com/gussev/d1/core/vicin"
	"github.com/gussev/d1/object"
)
// Builder has a function used to build the object
type Builder interface{
	/*
		title   - name of the object
		style   - subtype of the object
	    v       - environment parameters read from a configuration file
	    context - goroutine dependant information
	    object  - digital representation of the picture the program should detect
	    error   - in case any
	*/
	Build(title string,style string,v vicin.Vicinity, context context.Context) (object.Object,error)
}
