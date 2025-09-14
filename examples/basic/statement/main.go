package main

import "fmt"

type user struct {
	name    string
	age     int
	hobbies []string
}

// if else 文の代わりに switch
func main() {
	u := user{
		name:    "Taro",
		age:     21,
		hobbies: []string{"Music", "Sing a song"},
	}

	switch {
	case u.name == "":
		fmt.Println("Unknown")
	case u.age > 20:
		fmt.Println("Good")
	case len(u.hobbies) > 0:
		fmt.Println("Have fun")
	}
}
