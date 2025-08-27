package datetimeutils

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// システム日付をmockする
func withMockNow(mock time.Time, fn func()) {
	org := NowFunc
	NowFunc = func() time.Time { return mock }
	defer func() { NowFunc = org }()
	fn()
}

func TestNowUTC_WithMock(t *testing.T) {
	assert := assert.New(t)
	mock := time.Date(2025, 8, 23, 12, 34, 56, 0, time.UTC)
	withMockNow(mock, func() {
		expected := mock.UTC()
		actual := NowUTC()
		assert.Equal(expected, actual)
		fmt.Printf("expected=%v, actual=%v", expected, actual)
	})
}

func TestNowJST_WithMock(t *testing.T) {
	assert := assert.New(t)
	mock := time.Date(2025, 8, 23, 0, 0, 0, 0, time.UTC)
	withMockNow(mock, func() {
		expected := mock.In(JST)
		actual := NowJST()
		assert.Equal(expected, actual)
		fmt.Printf("expected=%v, actual=%v", expected, actual)
	})
}

func TestUTCToJST(t *testing.T) {
	assert := assert.New(t)
	expected := time.Date(2025, 8, 23, 9, 0, 0, 0, JST)
	actual := UTCToJST(time.Date(2025, 8, 23, 0, 0, 0, 0, time.UTC))
	assert.Equal(expected, actual)
	fmt.Printf("expected=%v, actual=%v", expected, actual)
}

func TestAddDaysEndOfDay(t *testing.T) {
	assert := assert.New(t)
	expected := time.Date(2025, 1, 31, 23, 59, 59, 0, JST)
	actual := AddDaysEndOfDay(time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC), 30)
	assert.Equal(expected, actual)
	fmt.Printf("expected=%v, actual=%v", expected, actual)
}

func TestAddDaysDaysStartOfDay(t *testing.T) {
	assert := assert.New(t)
	expected := time.Date(2025, 2, 1, 0, 0, 0, 0, JST)
	actual := AddDaysStartOfDay(time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC), 31)
	assert.Equal(expected, actual)
	fmt.Printf("expected=%v, actual=%v", expected, actual)
}

func TestNowMinusDaysStartOfDay_WithMock(t *testing.T) {
	assert := assert.New(t)
	mock := time.Date(2025, 8, 23, 15, 45, 0, 0, time.UTC)
	withMockNow(mock, func() {
		actual := NowMinusDaysStartOfDay(30)
		expected := time.Date(2025, 7, 24, 0, 0, 0, 0, JST)
		assert.Equal(expected, actual)
		fmt.Printf("expected=%v, actual=%v", expected, actual)
	})
}

func TestNowPlusDaysMinus1StartOfDay_WithMock(t *testing.T) {
	assert := assert.New(t)
	mock := time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC)
	withMockNow(mock, func() {
		actual := NowPlusDaysStartOfDay(30)
		expected := time.Date(2025, 1, 31, 0, 0, 0, 0, JST)
		assert.Equal(expected, actual)
		fmt.Printf("expected=%v, actual=%v", expected, actual)
	})
}
