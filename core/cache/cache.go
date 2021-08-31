package cache

import (
	"github.com/gussev/d1/object"
)

// Cache functions to work with local cache
type Cache interface {
	Put(key string,obj object.Object )
	Get(key string) object.Object
	Len() int
}
