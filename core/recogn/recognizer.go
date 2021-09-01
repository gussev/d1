package recogn

import (
	"github.com/gussev/d1/core/bitset"
	"github.com/gussev/d1/core/vicin"
)

// Recognizer to recognize simplified representation
type Recognizer interface{
	/*
	bs - simplified object presentation
	vicinity - environment data
	*/
	Recognize(bs bitset.BitSet,v vicin.Vicinity) string
}
