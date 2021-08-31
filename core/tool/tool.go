package tool

import (
	"image"
)
// Rectangle set of tools I could not find in Go, it seems will be deprecated
type Rectangle interface{
	String() string
	ImageRect() image.Rectangle
}
