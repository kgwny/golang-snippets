package main

import "fmt"

func main() {
	iNums := []int{1, 2, 3}
	var oNums []*int

	// ループ変数nのアドレスをoNumsへ追加する
	for _, n := range iNums {
		oNums = append(oNums, &n)
	}

	// oNumsに追加された値(index, value)およびアドレスを確認する
	for i, pn := range oNums {
		fmt.Printf("index [%d], value [%d], address [%p]\n", i, *pn, pn)
	}

}

// 実行結果
// index [0], value [1], address [0x14000092020]
// index [1], value [2], address [0x14000092028]
// index [2], value [3], address [0x14000092030]
