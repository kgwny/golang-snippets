package main

import "fmt"

// 整数定数値にはString()メソッドを定義しておこう
//  (Add String() method for integers const values)

type State int

// 定数
const (
	Running State = iota
	Stopped
	Rebooting
	Terminated
)

// State 型に対する String メソッド（レシーバを持つメソッド定義）
func (s State) String() string {
	switch s {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	case Rebooting:
		return "Rebooting"
	case Terminated:
		return "Terminated"
	default:
		return "Unknown"
	}
}

func main() {
	state := Running
	// レシーバを持つメソッド定義がない場合は、0が表示される
	fmt.Println(state)
}
