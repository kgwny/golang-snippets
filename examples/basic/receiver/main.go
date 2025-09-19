package main

import "fmt"

// 構造体
type Rectangle struct {
	Width, Height int
}

// 値レシーバー(値レシーバー → 対象オブジェクトのコピーに対して処理)
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// ポインタレシーバー(ポインタレシーバー → 元のオブジェクトそのものを変更)
func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	rect := Rectangle{2, 3}
	fmt.Println("面積:", rect.Area()) // 6

	rect.Scale(2)             // ポインタレシーバーなので元の変数を変更
	fmt.Println("拡大後:", rect) // {4 6}
}
