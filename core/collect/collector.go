package collect

import (
	"github.com/gussev/d1/core/bitset"
	"github.com/gussev/d1/core/stat"
	"github.com/gussev/d1/core/vicin"
	"github.com/gussev/d1/object"
)


// Collector set of functions to process gathered statistical information
type Collector interface{
	// Add new statistics to the collector
	Add(key string,stat stat.Statistics)
	// Compute gathered data
	Compute()
	// Log give out computed data
	Log(v vicin.Vicinity)
	// Statistics process stat data for current Object
	Statistics(bs bitset.BitSet,obj object.Object,v vicin.Vicinity,title string,style string) error
	// Data give out StatData
	Data() StatData
	// Take a strategy to pick up a right Object
	Take(t Taker) string
}
