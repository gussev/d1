package object

import (
	"github.com/gussev/d1/core/bitset"
	"github.com/gussev/d1/core/stat"
)

// Object represents a digital image of the picture the program should detect
type Object interface {
	Compute(bitset bitset.BitSet, s stat.Statistics) error
}

