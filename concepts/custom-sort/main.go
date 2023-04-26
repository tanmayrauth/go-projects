package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (b byLength) Len() int {
	return len(b)
}

func (b byLength) Swap(x, y int) {
	b[x], b[y] = b[y], b[x]
}

func (b byLength) Less(i, j int) bool {
	return len(b[i]) < len(b[j])
}

func main() {

	fruits := []string{"mango", "banana", "orange"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

}
