package tool

import (
	"image"
)
// Rectangle here will be something I have not found in Go yet.
type Rectangle interface{
	String() string
	ImageRect() image.Rectangle
}
