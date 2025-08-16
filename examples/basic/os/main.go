package main

import (
	"fmt"
	"os"
)

func main() {
	homeDir, _ := os.UserHomeDir()
	fmt.Println("Your home directory is:", homeDir)
}
