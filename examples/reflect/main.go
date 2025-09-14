package main

import (
	"fmt"
	"reflect"
)

type Sender interface {
	Send()
}

type Receiver interface {
	Recv()
}

type App struct{}

// App は Sender インターフェースだけを実装する
// Receiver インターフェースは実装しない
func (c *App) Send() {
	fmt.Println("Send something")
}

// ある変数がインターフェースを実装しているかチェックする
// reflect パッケージの Implements を用いる
func main() {
	a := &App{}

	at := reflect.TypeOf(a)
	st := reflect.TypeOf((*Sender)(nil)).Elem()
	rt := reflect.TypeOf((*Receiver)(nil)).Elem()

	fmt.Printf("Client implements Sender: %v\n", at.Implements(st))
	fmt.Printf("Client implements Receiver: %v\n", at.Implements(rt))
}

// Client implements Sender: true
// Client implements Receiver: false
