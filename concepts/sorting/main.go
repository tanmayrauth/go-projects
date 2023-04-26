package main

import (
	"fmt"
	"sort"
)

func main() {

	ints := []int{2, 3, 5} // It must always be slice, array will create an issue here.
	sort.Ints(ints)
	fmt.Println(ints)

	strs := []string{"a", "k", "d"}
	sort.Strings(strs)
	fmt.Println(strs)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Is Sorted:", s)

}
