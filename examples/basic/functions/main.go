package main

import "fmt"

// 関数の書き方
// func <関数名>([引数]) [戻り値の型] {
// 	   [関数の本体]
// }

func greetings(arg string) string {
	return arg
}

func add(x, y int) int {
	return x + y
}

// 返却値が複数あるパターン
func multipleArgs(arg1, arg2 string) (string, string) {
	return arg2, arg1
}

func main() {
	fmt.Println(greetings("Hello World!")) // Hello World!
	fmt.Println(add(1, 2))                 // 3
	a, b := multipleArgs("hoge", "fuga")
	fmt.Println(a, b) // fuga hoge
}
