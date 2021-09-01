package cache

import (
	"github.com/gussev/d1/object"
)

// Cache functions to work with local cache
type Cache interface {
	// Put object in Cache
	Put(key string,obj object.Object )
	// Get object from Cache
	Get(key string) object.Object
	// Len get amount elements in Cache
	Len() int
}
