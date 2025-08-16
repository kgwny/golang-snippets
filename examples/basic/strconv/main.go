package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "123"
	num, _ := strconv.Atoi(numStr)
	fmt.Println(num + 456) // 579
}
