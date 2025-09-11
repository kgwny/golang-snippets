package main

import "fmt"

// make していない map へのアクセス

type user struct {
	name string
	age  int
}

func main() {
	var m map[string]user

	// m は nil の状態
	if m == nil {
		fmt.Println("m is nil")
	}

	// キーにアクセスすると値として定義した型（user）の初期値が返る
	u := m["AAA"]
	fmt.Printf("%#v\n", u)

	// 初期化した場合
	m = make(map[string]user)

	// key に value をセットする
	m["BBB"] = user{name: "Taro", age: 18}
	fmt.Printf("%#v\n", m)
}
