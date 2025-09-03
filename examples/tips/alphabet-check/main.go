package main

import (
	"fmt"
	"regexp"
)

func main() {
	// チェック対象の文字列配列
	words := []string{"Hello", "World", "Go123", "ABC", "2025-09-04"}

	// アルファベットのみを許可する正規表現
	re := regexp.MustCompile(`^[A-Za-z]+$`)

	for _, word := range words {
		if re.MatchString(word) {
			fmt.Printf("%q はアルファベットのみで構成されています\n", word)
		} else {
			fmt.Printf("%q はアルファベット以外の文字を含みます\n", word)
		}

	}
}
