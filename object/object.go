package object

import (
	"github.com/gussev/d1/core/bitset"
	"github.com/gussev/d1/core/stat"
)

// Object represents a digital image of the picture the program should detect
type Object interface {
	// Compute statistical information from the current Object
	Compute(bitset bitset.BitSet, s stat.Statistics) error
	Key() Object_Key
}

type Object_Key interface {
	Mine(other Object_Key)bool
}

