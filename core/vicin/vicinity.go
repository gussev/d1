package vicin

import (
	"image"
	"image/color"
	"github.com/gussev/d1/core/cache"
	"github.com/gussev/d1/core/tool"
)

// Vicinity represents environment parameters read from a configuration file
type Vicinity interface{
	Put(key string,value interface{} )
	Rectangle(key string) tool.Rectangle
	// Reader cycle dependencies are not allowed in Go, instead of interface{} should be Reader
	Reader() interface{}
	Cache() cache.Cache
	// Recognizer cycle dependencies are not allowed in Go, instead of interface{} should be Recognizer
	Recognizer() interface{}
	Color(key string) color.RGBA
	String(key string) string
	Int(key string) int
	Bool(key string) bool
	StrArray(key string) []string
	Rect() image.Rectangle
	File(title string,style string, file_pattern string) string
	Key(title string,style string) string
	App_type(t string) bool
}