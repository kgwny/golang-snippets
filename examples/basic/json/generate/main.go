package main

import (
	"encoding/json"
	"fmt"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		Id:      1,
		Content: "greetings",
		Author: Author{
			Id:   1,
			Name: "Taro",
		},
		Comments: []Comment{
			{
				Id:      1,
				Content: "hoge",
				Author:  "Jiro",
			},
			{
				Id:      2,
				Content: "fuga",
				Author:  "Saburo",
			},
			{
				Id:      3,
				Content: "piyo",
				Author:  "Hanako",
			},
		},
	}

	// 出力結果にインデントなし
	out1, err := json.Marshal(&post)
	if err != nil {
		fmt.Println("error marshalling to json.", err)
		return
	}
	fmt.Println(string(out1))

	// 出力結果にインデントあり
	out2, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		fmt.Println("error marshalling with indent to json.", err)
		return
	}
	fmt.Println(string(out2))
}
