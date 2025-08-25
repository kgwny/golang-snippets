package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	for i := range 10 {
		fmt.Println(i+1, ":", rand.N(10*time.Minute))
	}
}
