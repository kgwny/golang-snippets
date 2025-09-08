package main

import "fmt"

func main() {
	// 基本的なfor文
	fmt.Println("1.基本的なfor文")
	for i := 0; i < 5; i++ {
		fmt.Println("i = ", i)
	}

	// while文のように使う
	fmt.Println("\n2.while風のfor文")
	j := 0
	for j < 5 {
		fmt.Println("j = ", j)
		j++
	}

	// 無限ループ
	fmt.Println("\n3.無限ループ（3回でbreak）")
	k := 0
	for {
		if k >= 3 {
			break
		}
		fmt.Println("k =", k)
		k++
	}

	// rangeを使って配列やスライスをループ
	fmt.Println("\n4.rangeを使ったループ")
	numbers := []int{10, 20, 30, 40, 50}
	for idx, value := range numbers {
		fmt.Printf("index=%d, value=%d\n", idx, value)
	}

	// mapをループ
	fmt.Println("\n5.mapをループ")
	fruits := map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"grape":  "purple",
	}
	for key, val := range fruits {
		fmt.Printf("%s -> %s\n", key, val)
	}
}
