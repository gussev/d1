package collect

import (
	"github.com/gussev/d1/core/bitset"
	"github.com/gussev/d1/core/stat"
	"github.com/gussev/d1/core/vicin"
	"github.com/gussev/d1/object"
)


// Collector set of functions to process gathered statistical information
type Collector interface{
	Add(key string,stat stat.Statistics)
	Compute()
	Log(v vicin.Vicinity)
	Statistics(bs bitset.BitSet,obj object.Object,v vicin.Vicinity,title string,style string) error
	Data() StatData
	Take(t Taker) string
}
