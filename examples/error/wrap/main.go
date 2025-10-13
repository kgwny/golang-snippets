package main

import (
	"errors"
	"fmt"
)

func readFile() error {
	return errors.New("file not found")
}

func process() error {
	err := readFile()
	if err != nil {
		// %w を用いることで、エラーをラップして上位へ伝搬する
		return fmt.Errorf("process failed: %w", err)
	}
	return nil
}

func main() {
	err := process()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
