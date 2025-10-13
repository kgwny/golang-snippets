package main

import "fmt"

// `defer` + `recover()` によるパニック処理（例外的ケース）
// Go では `panic` が発生した場合、`recover` を使って
// プログラムのクラッシュを防ぐことができる
// ただし、**通常のエラー処理には使わない** のが原則

func risky() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("something went wrong")
}

func main() {
	risky()
	fmt.Println("Program continues")
}
