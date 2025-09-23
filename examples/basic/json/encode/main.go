package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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
	jsonFile, err := os.Open("test.json")
	if err != nil {
		fmt.Println("can't open json file.", err)
	}

	defer jsonFile.Close()
	// ioutil.ReadAll は deprecated. 代わりに io.ReadAll を使用する
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("unable to load json data.", err)
		return
	}

	var post Post
	json.Unmarshal(jsonData, &post)

	fmt.Println(post)
	fmt.Println(post.Comments)
	fmt.Println(post.Comments[0].Content)
}
