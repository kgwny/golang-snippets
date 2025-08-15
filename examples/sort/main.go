package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int{65535, 222, 123, 999, 234}
	sort.Ints(values)
	fmt.Println(values) // [123 222 234 999 65535]
}
