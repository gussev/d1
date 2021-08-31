package bitset

// BitSet functions to work with simplified picture
type BitSet interface{
	Get(index int) bool
	Set(index int, value bool)
	Len() int
}



