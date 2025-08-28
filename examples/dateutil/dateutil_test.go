package dateutil

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNowJST(t *testing.T) {
	// モックする基準時刻
	mockNow := time.Date(2025, 8, 24, 12, 34, 56, 0, time.UTC)

	// NowFunc を差し替える
	NowFunc = func() time.Time { return mockNow }
	// 差し替えた NowFunc を元に戻す
	defer func() { NowFunc = time.Now }()

	actual := NowJST()
	expected := mockNow.In(JST)

	assert.Equal(t, expected, actual)
	fmt.Printf("expected=%v, actual=%v", expected, actual)
}

func TestAddDaysSartOfDay(t *testing.T) {
	base := time.Date(2025, 8, 24, 15, 30, 0, 0, JST)
	actual := AddDaysStartOfDay(base, 2)
	expected := time.Date(2025, 8, 26, 0, 0, 0, 0, JST)
	assert.Equal(t, expected, actual)
	fmt.Printf("expected=%v, actual=%v", expected, actual)
}

func TestAddDaysEndOfDay(t *testing.T) {
	base := time.Date(2025, 8, 24, 15, 30, 0, 0, JST)
	actual := AddDaysEndOfDay(base, -1)
	expected := time.Date(2025, 8, 23, 23, 59, 59, 0, JST)
	assert.Equal(t, expected, actual)
	fmt.Printf("expected=%v, actual=%v", expected, actual)
}

func TestNowMinusDaysStartOfDay(t *testing.T) {
	// 2025-08-24 10:00:00 JST を基準にする
	mockNow := time.Date(2025, 8, 24, 10, 0, 0, 0, JST)
	NowFunc = func() time.Time { return mockNow }
	defer func() { NowFunc = time.Now }()

	actual := NowMinusDaysStartOfDay(3) // 3日前
	expected := time.Date(2025, 8, 21, 0, 0, 0, 0, JST)
	assert.Equal(t, expected, actual)
	fmt.Printf("expected=%v, actual=%v", expected, actual)
}

func TestNowPlusDaysStartOfDay(t *testing.T) {
	// 2025-08-24 10:00:00 JST を基準にする
	mockNow := time.Date(2025, 8, 24, 10, 0, 0, 0, JST)
	NowFunc = func() time.Time { return mockNow }
	defer func() { NowFunc = time.Now }()

	actual := NowPlusDaysStartOfDay(5) // 5日後
	expected := time.Date(2025, 8, 29, 0, 0, 0, 0, JST)
	assert.Equal(t, expected, actual)
	fmt.Printf("expected=%v, actual=%v", expected, actual)
}
