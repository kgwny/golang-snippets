package main

import (
	"fmt"
	"time"
)

// 日時取得
func main() {
	// 現在の時刻（UTC基準で取得する）
	now := time.Now()

	// 表示フォーマット
	// 日付・時刻のフォーマットは、GO特有の基準日 "2006-01-02 15:04:05" を使って指定する
	const layout = "2006-01-02 15:04:05"

	// UTC
	utc := now.UTC().Format(layout)
	fmt.Println("UTC:", utc)

	// JST (UTC+9)
	jst := now.In(time.FixedZone("JST", 9*60*60)).Format(layout)
	fmt.Println("JST:", jst)
}
