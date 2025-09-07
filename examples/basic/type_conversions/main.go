package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 整数型の変換
	var i int = 42
	var i8 int8 = int8(i)
	var i16 int16 = int16(i)
	var i32 int32 = int32(i)
	var i64 int64 = int64(i)
	var u uint = uint(i)
	var u8 uint8 = uint8(i)
	var u16 uint16 = uint16(i)
	var u32 uint32 = uint32(i)
	var u64 uint64 = uint64(i)

	fmt.Println("整数型変換:")
	fmt.Println(i, i8, i16, i32, i64, u, u8, u16, u32, u64)

	// 浮動小数点型の変換
	var f32 float32 = float32(i)
	var f64 float64 = float64(i)
	fmt.Println("\n整数 -> 浮動小数点:", f32, f64)

	// 浮動小数点 -> 整数（小数点以下切り捨て）
	var f float64 = 3.14
	var fi int = int(f)
	fmt.Println("\n浮動小数点 -> 整数:", fi)

	// rune (Unicode code point) と byte (uint8)
	var ch rune = 'A' // rune は int32 の別名
	var b byte = 'B'  // byte はuint8の別名
	fmt.Println("\ntune と byte:", ch, string(ch), b, string(b))

	// 文字列と数値の変換
	var str string = "123"
	num, _ := strconv.Atoi(str)   // string -> int
	strFromInt := strconv.Itoa(i) // int -> string
	fmt.Println("\n文字列と数値変換:", str, num, strFromInt)

	// string と float
	strFloat := "3.1415"
	fnum, _ := strconv.ParseFloat(strFloat, 64) // string -> float64
	strFromFloat := fmt.Sprintf("%f", fnum)     // float -> string
	fmt.Println("\n文字列と不動小数点変換:", strFloat, fnum, strFromFloat)

	// bool と string
	boolStr := "true"
	bv, _ := strconv.ParseBool(boolStr) // string -> bool
	strFromBool := strconv.FormatBool(bv)
	fmt.Println("\n文字列とbool変換:", boolStr, bv, strFromBool)

	// 整数と文字列（Format/Parse）
	strHex := strconv.FormatInt(int64(i), 16) // int -> 16進文字列
	parsedInt, _ := strconv.ParseInt(strHex, 16, 64)
	fmt.Println("\n整数と16進文字列変換:", strHex, parsedInt)

}
