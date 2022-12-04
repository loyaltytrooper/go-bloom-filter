package database

type DB struct {
	Size   int
	Values []interface{}
}

// whenever the number of 1's become 70% of size in BloomFilterCache
// we have to re-allocate all of the storage by dynamically increasing the size of
// the object used

type BloomFilterCache struct {
	Size int
	Bits []int
}
