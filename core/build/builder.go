package build

import (
	"github.com/gussev/d1/core/context"
	"github.com/gussev/d1/core/vicin"
	"github.com/gussev/d1/object"
)
// Builder has a function used to build the object
type Builder interface{
	/*
		title - name of the object
		style - subtype of the object
	*/
	Build(title string,style string,v vicin.Vicinity, context context.Context) (object.Object,error)
}
