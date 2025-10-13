package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func findUser(id int) error {
	if id == 0 {
		return ErrNotFound
	}
	return nil
}

func main() {
	err := findUser(0)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("User not found.")
	}
}
