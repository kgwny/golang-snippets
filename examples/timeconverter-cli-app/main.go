package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const layout = "2006-01-02 15:04:05"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== JST <-> UTC Time Converter ===")
	fmt.Println("1: JST -> UTC")
	fmt.Println("2: UTC -> JST")
	fmt.Print("Choose option (1 or 2): ")

	opt, _ := reader.ReadString('\n')
	opt = strings.TrimSpace(opt)

	switch opt {
	case "1":
		fmt.Print("Enter JST datetime (e.g., 2025-08-03 10:00:00): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			fmt.Println("Failed to load JST timezone:", err)
			return
		}

		t, err := time.ParseInLocation(layout, input, jst)
		if err != nil {
			fmt.Println("Parse error:", err)
			return
		}

		fmt.Println("UTC:", t.UTC().Format(layout))

	case "2":
		fmt.Print("Enter UTC datetime (e.g., 2025-08-03 01:00:00): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		t, err := time.Parse(layout, input)
		if err != nil {
			fmt.Println("Parse error:", err)
			return
		}

		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			fmt.Println("Failed to load JST timezone:", err)
			return
		}

		fmt.Println("JST:", t.In(jst).Format(layout))

	default:
		fmt.Println("Invalid option.")
	}
}
