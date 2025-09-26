package main

import (
	"fmt"
)

// interface の定義
type Animal interface {
	Speak() string
}

// Dog 型
type Dog struct {
	Name string
}

// Dog が Animal インターフェースを実装
func (d Dog) Speak() string {
	return d.Name + " says: Woof!"
}

// Cat 型
type Cat struct {
	Name string
}

// Cat が Animal インターフェースを実装
func (c Cat) Speak() string {
	return c.Name + " says: Meow!"
}

// Animal のスライスを受け取って動作を実行する関数
func MakeNoise(animals []Animal) {
	for _, a := range animals {
		fmt.Println(a.Speak())
	}
}

func main() {
	// Animal 型として Dog と Cat を使う
	animals := []Animal{
		Dog{Name: "Buddy"},
		Cat{Name: "Kitty"},
	}

	MakeNoise(animals)
}
