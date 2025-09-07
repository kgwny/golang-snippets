package main

import "fmt"

// 関数の外では := を利用した変数宣言は不可
var str = "hoge fuga piyo"
var num = 2
var var1, var2 bool

var s, n, b = "foo", 8, true

func main() {
	fmt.Println(str, num, var1, var2) // hoge fuga piyo 2 false false
	fmt.Println(s, n, b)              // foo, 8, true

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("初期値: %v, %v, %v, %q\n", i, f, b, s) // 初期値: 0, 0, false, ""
}

// 型
// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte = uint8
// rune = int32
// float32 float64
// complex64 complex128
