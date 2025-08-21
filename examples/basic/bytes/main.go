package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.NewBufferString("Golang is awesome!")
	fmt.Println(buffer.String())
}
