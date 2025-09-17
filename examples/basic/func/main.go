package main

import "fmt"

// func 関数名(引数リスト) 戻り値型 {
// 	// 処理
// }

// 関数の基本構文
func add(a int, b int) int {
	return a + b
}

// 複数引数で同じ型の関数
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 複数戻り値の関数
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// 名前付き戻り値(関数の戻り値に変数名を付与する)
func rect(w, h int) (area int, perimeter int) {
	area = w * h
	perimeter = 2 * (w + h)
	return
}

// 関数を引数に渡すパターン
// Goでは関数が「第一級オブジェクト」なので、関数を渡したり返却したりできる
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// 関数を戻り値にするパターン
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// メソッド(レシーバー付き関数)
// funcは型に紐づけてメソッドとして定義することもできる
type Rectangle struct {
	Width, Height int
}

// 値レシーバー
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// ポインタレシーバー
func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	// 無名関数(関数リテラル)
	// 関数を変数に代入したり、その場で定義して実行することができる
	f := func(a, b int) int {
		return a + b
	}
	fmt.Println(f(3, 4)) // 7

	// 無名関数の即時実行パターン
	result := func(x int) int {
		return x * x
	}(5)
	fmt.Println(result) // 25

	// add関数
	c := add(2, 3)
	fmt.Println("2 + 3 = ", c) // 5

	// sum関数
	sum := sum(1, 2, 3, 4, 5)
	fmt.Println("sum = ", sum) // 15

	// divide関数
	div, _ := divide(6, 2)
	fmt.Println("div = ", div) // 3

	// rect関数
	area, perimeter := rect(2, 4)
	fmt.Println("area = ", area)
	fmt.Println("perimeter = ", perimeter)

	// operate関数
	res := operate(44, 55, func(x, y int) int {
		return x + y
	})
	fmt.Println(res) // 99

	// multiplier関数
	double := multiplier(2) // factor = 2
	triple := multiplier(3) // factor = 3
	fmt.Println(double(5))  // 10 (5 * 2)
	fmt.Println(triple(5))  // 15 (5 * 3)

	// Area関数、Scale関数
	rect := Rectangle{Width: 3, Height: 4}
	fmt.Println(rect.Area()) // 12

	// ポインタを明示して呼び出し
	(&rect).Scale(2)
	// rect自体が変更される
	fmt.Println(rect.Width, rect.Height) // 6 8

	fmt.Println("rect = ", rect)
}
