package main

import (
	"fmt"
	"time"
)

// 日付取得
func main() {
	// 日本時間（JST）のロケーションを取得
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	// 現在の時刻をJSTで取得
	now := time.Now().In(loc)

	// 日付だけをフォーマットして表示
	fmt.Println(now.Format("2006-01-02"))
}
