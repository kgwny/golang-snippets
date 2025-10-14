package main

import (
	"fmt"
	"math"
)

const s string = "hoge"

func main() {

	const n = 600000000
	const d = 3e20 / n

	fmt.Println("s =", s)
	fmt.Println("d =", d)
	fmt.Println("int64(d) =", int64(d))
	fmt.Println("math.Sin(n) =", math.Sin(n))
}
