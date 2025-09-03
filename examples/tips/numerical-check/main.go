package main

import (
	"fmt"
	"strings"
)

// 文字列が数字のみで構成されていることをチェックする
func main() {
	params := []string{"0", "hoge", "m2", "2025-09-03", "65535"}

	for _, v := range params {
		if strings.Trim(v, "0123456789") == "" {
			fmt.Println(v)
		}
	}
}
