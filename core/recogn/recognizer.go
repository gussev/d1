package recogn

import (
	"github.com/gussev/d1/core/bitset"
	"github.com/gussev/d1/core/vicin"
)

// Recognizer to recognize simplified representation
type Recognizer interface{
	Recognize(bs bitset.BitSet,v vicin.Vicinity) string
}
