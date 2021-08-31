package collect

import (
	"github.com/gussev/d1/core/stat"
	"github.com/gussev/d1/core/vicin"
)
// StatData set of helper functions to process statistical information
type StatData interface{
	Add(key string,stat stat.Statistics)
	Compute()
	Log(v vicin.Vicinity)
	Len() int
	Select() (string,interface{})
}
