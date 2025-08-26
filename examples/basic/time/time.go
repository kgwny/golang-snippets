package time

import (
	"fmt"
	"time"
)

// 日時取得
func main() {
	// 現在の時刻（UTC基準で取得する）
	now := time.Now()

	// UTC と JST のフォーマット済み文字列を取得する
	utc, jst := GetFormattedTimes(now)

	// 現在時刻
	fmt.Println("now:", now)

	// UTC
	fmt.Println("UTC:", utc)

	// JST (UTC+9)
	fmt.Println("JST:", jst)
}

func GetFormattedTimes(t time.Time) (string, string) {
	// 表示フォーマット
	// 日付・時刻のフォーマットは、GO特有の基準日 "2006-01-02 15:04:05" を使って指定する
	const layout = "2006-01-02 15:04:05"

	utc := t.UTC().Format(layout)
	jst := t.In(time.FixedZone("JST", 9*60*60)).Format(layout)

	return utc, jst
}
