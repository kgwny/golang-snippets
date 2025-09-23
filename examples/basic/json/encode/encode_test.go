package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// go test -v で実行
func TestLoadPostFromJSONFile(t *testing.T) {
	// テスト用データ
	post := Post{
		Id:      1,
		Content: "hello",
		Author:  Author{Id: 1, Name: "Taro"},
		Comments: []Comment{
			{Id: 1, Content: "hoge", Author: "Jiro"},
			{Id: 2, Content: "fuga", Author: "Saburo"},
		},
	}

	// JSON にシリアライズ
	data, err := json.Marshal(post)
	if err != nil {
		t.Fatalf("failed to marshal post: %v", err)
	}

	// 一時ディレクトリに test.json を作成
	tmpDir := t.TempDir()
	jsonPath := filepath.Join(tmpDir, "test.json")
	if err := os.WriteFile(jsonPath, data, 0644); err != nil {
		t.Fatalf("failed to write test.json: %v", err)
	}

	// ファイルを読み込んで Unmarshal
	content, err := os.ReadFile(jsonPath)
	if err != nil {
		t.Fatalf("failed to read test.json: %v", err)
	}

	var loaded Post
	if err := json.Unmarshal(content, &loaded); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	// 内容チェック
	if loaded.Content != "hello" {
		t.Errorf("expected content %q, got %q", "hello", loaded.Content)
	}
	if loaded.Author.Name != "Taro" {
		t.Errorf("expected author name %q, got %q", "Taro", loaded.Author.Name)
	}
	if len(loaded.Comments) != 2 {
		t.Fatalf("expected 2 comments, got %d", len(loaded.Comments))
	}
	if loaded.Comments[0].Content != "hoge" {
		t.Errorf("expected first comment content %q, got %q", "hoge", loaded.Comments[0].Content)
	}
	if loaded.Comments[1].Author != "Saburo" {
		t.Errorf("expected second comment author %q, got %q", "Saburo", loaded.Comments[1].Author)
	}
}
