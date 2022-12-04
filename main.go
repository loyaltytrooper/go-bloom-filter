package main

import (
	"bloom-filters/database"
	"bloom-filters/service"
	"fmt"
)

func main() {
	cache := database.BloomFilterCache{}
	cache.OnBits = 0
	cache.BitsArray = make([]byte, 10)

	database := database.DB{}
	var rehash = true

	for rehash == true {
		rehash = false
		input := []string{"Hi", "this is ", "my ", "internship ", "assignment"}
		service.SetCacheBits(input, &database, &cache, &rehash)

		if rehash == true {
			cache.BitsArray = make([]byte, len(cache.BitsArray)*2)
		}
	}

	fmt.Println("Completed setting values in BloomFilter Cache\n")

	checkArray := []string{"Wassup", "this is ", "my ", "internship ", "you", "assigned"}
	service.CheckInCache(checkArray, &cache, &database)
}
