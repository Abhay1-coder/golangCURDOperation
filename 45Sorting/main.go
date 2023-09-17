// main.go
package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"c", "a", "b"}
	slices.SortStrings(strs)
	fmt.Println("Sorted Strings:", strs)

	ints := []int{7, 2, 4}
	slices.SortInts(ints)
	fmt.Println("Sorted Ints:   ", ints)
}
