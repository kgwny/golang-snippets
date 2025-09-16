package main

import (
	"fmt"

	"github.com/google/uuid"
)

// uuidのインストール
// go get -u github.com/google/uuid

func main() {
	// UUID v4 を生成する
	u := uuid.New()
	fmt.Println("UUID v4:", u.String())

	// 文字列からUUIDをパースしたい場合
	s := "550e8400-e29b-41d4-a716-446655440000"
	parsed, err := uuid.Parse(s)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}
	fmt.Println("Parsed UUID:", parsed)
}
