package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// 基本的なエラーハンドリング (`if err != nil`)
// 関数は通常 `(値, error)` の2つを返す
// 呼び出し側で `if err != nil` で確認
// Go では try-catch がないため、このパターンが基本となる
func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}
