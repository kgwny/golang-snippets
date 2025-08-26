package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetFormattedTimes(t *testing.T) {
	baseTime := time.Date(2025, 1, 2, 3, 4, 5, 0, time.UTC)

	utc, jst := GetFormattedTimes(baseTime)

	assert.Equal(t, "2025-01-02 03:04:05", utc, "UTC フォーマットが正しいこと")
	assert.Equal(t, "2025-01-02 12:04:05", jst, "JST フォーマットが正しいこと")
}
