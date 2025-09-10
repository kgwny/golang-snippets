package main

import "fmt"

func main() {

	// スライスの初期化
	// make(配列の型, length, capacity)

	// スライスを作成する
	slice := make([]int, 3, 5)                 // 長さ3, 容量5のスライス
	fmt.Println(slice, len(slice), cap(slice)) // [0 0 0] 3 5

	// スライスに要素を追加する
	slice = append(slice, 4, 5)                // 容量内に追加
	fmt.Println(slice, len(slice), cap(slice)) // [0 0 0 4 5] 5 5

	// 容量を超過して要素を追加する
	slice = append(slice, 6)                   // 容量を超えてしまい、新しい基底配列が作成される
	fmt.Println(slice, len(slice), cap(slice)) // [0 0 0 4 5] 6 10
}

// [0 0 0] 3 5
// [0 0 0 4 5] 5 5
// [0 0 0 4 5 6] 6 10
