package main

import "fmt"

func main() {
	message := greetMe("world")
	fmt.Println(message)
}

func greetMe(name string) string {
	if name == "" {
		return "Hello!"
	}
	return "Hello, " + name + "!"
}
