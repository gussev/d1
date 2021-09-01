package bitset

// BitSet functions to work with simplified picture
type BitSet interface{
	// Get value of the index in bitset
	Get(index int) bool
	// Set value for the index in bitset
	Set(index int, value bool)
	// Len of the bitset
	Len() int
}



