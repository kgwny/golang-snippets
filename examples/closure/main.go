package main

import "fmt"

func getFunc() func() {
	i := 0

	// クロージャー
	return func() {
		i++
		fmt.Println(i)
	}
}

// 簡単なクロージャーの例
func main() {
	fmt.Println("fn1 start")
	fn1 := getFunc()
	fn1()
	fn1()
	fn1()
	fmt.Println("fn1 end")

	fmt.Println("fn2 start")
	fn2 := getFunc()
	fn2()
	fn2()
	fmt.Println("fn2 end")
}
