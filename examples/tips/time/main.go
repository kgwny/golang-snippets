package main

import (
	"fmt"
	"time"
)

// https://pkg.go.dev/time#pkg-constants
func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond())

	// ex) RFC3339 = "2006-01-02T15:04:05Z07:00"
	// 2025-09-20T14:14:36+09:00
	fmt.Println(t.Format(time.RFC3339))

	// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	// 2025-09-20T14:20:35.367048+09:00
	fmt.Println(t.Format(time.RFC3339Nano))

	// yyyy-MM-dd HH:mm:ss
	// 2025-09-22 10:26:47
	fmt.Println(t.Format(time.DateTime))

	// yyyy-MM-dd
	// 2025-09-22
	fmt.Println(t.Format(time.DateOnly))

	// HH:mm:ss
	// 10:26:47
	fmt.Println(t.Format(time.TimeOnly))
}
