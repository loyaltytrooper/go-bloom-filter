package main

import "bloom-filters/database"

func main() {
	cache := database.BloomFilterCache{}
	cache.Bits = make([]int, 100)
}
