package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("nonexistent.txt")
	if err != nil {
		var pathErr *os.PathError
		if ok := errors.As(err, &pathErr); ok {
			fmt.Println("Failed to open:", pathErr.Path)
		}
	}
}
