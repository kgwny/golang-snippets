package datetimeutils

import (
	"time"
)

// JSTのロケーション
var JST *time.Location

func init() {
	var err error
	JST, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic("failed to load Asia/Tokyo location: " + err.Error())
	}
}

// システム日付を取得する
var NowFunc = time.Now

// UTCで現在日時を取得する
func NowUTC() time.Time {
	return NowFunc().UTC()
}

// JSTで現在日時を取得する
func NowJST() time.Time {
	return NowFunc().In(JST)
}

// UTCの現在日時のタイムゾーンをJSTに変換する
func UTCToJST(t time.Time) time.Time {
	return t.In(JST)
}

// UTCの指定日付を基にして、N日後の 23:59:59 の日時を取得する
func AddDaysEndOfDay(t time.Time, days int) time.Time {
	target := t.AddDate(0, 0, days)
	return time.Date(target.Year(), target.Month(), target.Day(), 23, 59, 59, 0, JST)
}

// UTCの指定日付を基にして、N日後の 00:00:00 の日時を取得する
func AddDaysStartOfDay(t time.Time, days int) time.Time {
	target := t.AddDate(0, 0, days)
	return time.Date(target.Year(), target.Month(), target.Day(), 0, 0, 0, 0, JST)
}

// 現在日付のN日前の 00:00:00 の日時を取得する
func NowMinusDaysStartOfDay(days int) time.Time {
	now := NowFunc()
	target := now.AddDate(0, 0, -days)
	return time.Date(target.Year(), target.Month(), target.Day(), 0, 0, 0, 0, JST)
}

// 現在日付のN日後の 00:00:00 の日時を取得する
func NowPlusDaysStartOfDay(days int) time.Time {
	now := NowFunc()
	target := now.AddDate(0, 0, days)
	return time.Date(target.Year(), target.Month(), target.Day(), 0, 0, 0, 0, JST)
}
