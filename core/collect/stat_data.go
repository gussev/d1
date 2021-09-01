package collect

import (
	"github.com/gussev/d1/core/stat"
	"github.com/gussev/d1/core/vicin"
)
// StatData set of helper functions to process statistical information
type StatData interface{
	// Add stat data
	Add(key string,stat stat.Statistics)
	// Compute data
	Compute()
	// Log print the data
	Log(v vicin.Vicinity)
	// Len amount of collected stat data
	Len() int
	// Select the possible match
	Select() (string,interface{})
}
