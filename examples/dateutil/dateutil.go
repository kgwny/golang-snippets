package dateutil

import (
	"time"
)

// JST ロケーション（固定で利用する前提）
var JST *time.Location

func init() {
	var err error
	JST, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic("failed to load Asia/Tokyo location: " + err.Error())
	}
}

// システム日付を取得する（テスト時は差し替え可能）
var NowFunc = time.Now

func NowJST() time.Time {
	return NowFunc().In(JST)
}

func AddDaysStartOfDay(t time.Time, days int) time.Time {
	target := t.AddDate(0, 0, days)
	return time.Date(target.Year(), target.Month(), target.Day(), 0, 0, 0, 0, JST)
}

func AddDaysEndOfDay(t time.Time, days int) time.Time {
	target := t.AddDate(0, 0, days)
	return time.Date(target.Year(), target.Month(), target.Day(), 23, 59, 59, 0, JST)
}

func NowMinusDaysStartOfDay(days int) time.Time {
	now := NowFunc()
	target := now.AddDate(0, 0, -days)
	return time.Date(target.Year(), target.Month(), target.Day(), 0, 0, 0, 0, JST)
}

func NowPlusDaysStartOfDay(days int) time.Time {
	now := NowFunc()
	target := now.AddDate(0, 0, days)
	return time.Date(target.Year(), target.Month(), target.Day(), 0, 0, 0, 0, JST)
}
