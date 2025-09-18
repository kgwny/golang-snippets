package main

import "fmt"

// ポインターとは
// ポインターとは変数そのものではなく"変数が格納されているメモリアドレス"を
// 保持する特別な変数です。
// Go言語では * と & を用います。それぞれの持つ意味は以下のとおりです。
// & : 変数のアドレスを取得する
// * : アドレスから値を参照する（デリファレンス）

func add(a, b int) int {
	return a + b
}

func changeValue(n *int) {
	*n = *n * 2
}

type User struct {
	Name string
	Age  int
}

func updateAge(u *User, newAge int) {
	u.Age = newAge
}

func main() {
	// ポインタを使わないとき(当然ではあるが処理で扱う元の変数は変わらない)
	a := 5
	b := a * 2
	fmt.Println("a + b =", add(a, b))
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// 基本的な例
	x := 10 // 整数の変数
	p := &x // x のアドレスを格納するポインタ

	fmt.Println("xの値:", x) // 10
	fmt.Println("xのアドレス:", &x)
	fmt.Println("pの値:", p)    // &x と同じアドレス
	fmt.Println("pが指す値:", *p) // 10（デリファレンス）

	*p = 20                   // ポインタ経由で値を変更
	fmt.Println("xの新しい値:", x) // 20

	// 関数とポインタ
	// 値渡しの場合はコピーされるが、ポインタ渡しにすれば「元の変数」を操作できる
	y := 5
	changeValue(&y)
	fmt.Println(y) // 10

	// 構造体とポインタ
	user := User{Name: "Alice", Age: 25}
	updateAge(&user, 30)
	fmt.Println(user) // {Alice 30}

	// new と make
	// new(T) の使い方
	pp := new(int)
	fmt.Println(*pp) // 0
	*pp = 45
	fmt.Println(*pp) // 45

	// make(T, args) の使い方
	// スライス・マップ・チャネル専用で内部構造を初期化して返却する
	s := make([]int, 3)
	fmt.Println(s) // [0 0 0]
}
