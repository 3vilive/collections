package main

import (
	"fmt"

	"github.com/3vilive/collections/bitset"
)

func isAllAsciiCharsDiff(s string) bool {
	bitSet := bitset.NewBitSet(256)

	for _, char := range []byte(s) {
		if bitSet.Get(int(char)) {
			return false
		}

		bitSet.Set(int(char))
	}

	return true
}

func main() {
	inputArr := []string{
		"aAbBcC",
		"ABCDEF",
		"AABB",
		"ACBA",
	}

	for _, input := range inputArr {
		fmt.Printf("isAllAsciiCharsDiff(%s) = %v\n", input, isAllAsciiCharsDiff(input))
	}

	// isAllAsciiCharsDiff(aAbBcC) = true
	// isAllAsciiCharsDiff(ABCDEF) = true
	// isAllAsciiCharsDiff(AABB) = false
	// isAllAsciiCharsDiff(ACBA) = false
}
