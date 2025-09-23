package main

import (
	"encoding/json"
	"strings"
	"testing"
)

// go test -v で実行
func TestPostJSONMarshalling(t *testing.T) {
	post := Post{
		Id:      1,
		Content: "greetings",
		Author: Author{
			Id:   1,
			Name: "Taro",
		},
		Comments: []Comment{
			{Id: 1, Content: "hoge", Author: "Jiro"},
			{Id: 2, Content: "fuga", Author: "Saburo"},
			{Id: 3, Content: "piyo", Author: "Hanako"},
		},
	}

	// --- Marshal (no indent) ---
	out1, err := json.Marshal(&post)
	if err != nil {
		t.Fatalf("unexpected error in Marshal: %v", err)
	}
	jsonStr1 := string(out1)

	// 含まれるべき文字列の確認
	expectedKeywords := []string{
		"greetings", "Taro", "hoge", "fuga", "piyo",
	}
	for _, keyword := range expectedKeywords {
		if !strings.Contains(jsonStr1, keyword) {
			t.Errorf("expected JSON to contain %q, but got: %s", keyword, jsonStr1)
		}
	}

	// --- MarshalIndent (with indent) ---
	out2, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		t.Fatalf("unexpected error in MarshalIndent: %v", err)
	}
	jsonStr2 := string(out2)

	// インデントが含まれるかチェック（タブ文字）
	if !strings.Contains(jsonStr2, "\t\t") {
		t.Errorf("expected indented JSON to contain tabs, but got: %s", jsonStr2)
	}
}
