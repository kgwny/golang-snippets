package main

import (
	"fmt"
	"regexp"
)

/**
 * 正規表現
 *
 * 文字クラス（Character classes）
 * \d: 数字（0-9）
 * \D: 数字以外
 * \w: 単語文字（a-zA-Z0-9_）
 * \W: 単語文字以外
 * \s: 空白文字（スペース、タブ、改行など）
 * \S: 空白文字以外
 *
 * 量指定子（Quantifiers）
 * *: 直前の文字やグループが0回以上くりかえされる場合にマッチする
 * +: 直前の文字やグループが1回以上くりかえされる場合にマッチする
 * ?: 直前の文字やグループが0回または1回くりかえされる場合にマッチする
 * {n}: 直前の文字やグループがn回くりかえされる場合にマッチする
 * {n,}: 直前の文字やグループがn回以上くりかえされる場合にマッチする
 * {n,m}: 直前の文字やグループがn回以上、m回以下くりかえされる場合にマッチする
 *
 * 位置指定子（Anchors）
 * ^: 文字列の先頭
 * $: 文字列の末尾
 * \b: 単語の境界
 * \B: 単語の境界以外
 *
 * グループ（Groups）
 * (): キャプチャグループ（マッチした部分文字列を抽出する）
 * (?:): 非キャプチャグループ（マッチした部分文字列を抽出しない）
 */
func main() {
	text := "I am learning the Go language."

	// 正規表現パターン
	pattern := `\W`

	// コンパイル
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regexp:", err)
		return
	}

	// マッチング
	matches := re.FindAllString(text, -1)
	fmt.Println("Matches found:", matches)
}
