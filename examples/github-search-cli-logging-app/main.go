package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type GitHubResponse struct {
	Items []Repository `json:"items"`
}

type Repository struct {
	Name        string `json:"name"`
	HTMLURL     string `json:"html_url"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
}

// 検索キーワードを入力すると、GitHubの人気リポジトリ（スター数が多いもの）を最大10件取得して表示する/ログ出力するCLIアプリ
func main() {
	// ログファイル設定
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("ログファイルの作成に失敗しました:", err)
		return
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("アプリケーション開始")

	// キーワード入力
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("検索キーワードを入力してください: ")
	keyword, _ := reader.ReadString('\n')
	keyword = url.QueryEscape(keyword[:len(keyword)-1]) // 改行を除去してエスケープ
	count := 10

	log.Printf("検索キーワード: %s", keyword)

	apiURL := fmt.Sprintf("https://api.github.com/search/repositories?q=%s&sort=stars&order=desc&per_page=%d", keyword, count)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatalf("APIリクエスト失敗: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("GitHub API エラー: %s", resp.Status)
	}

	var result GitHubResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("JSONパース失敗: %v", err)
	}

	log.Printf("検索結果 %d 件取得", len(result.Items))

	// 結果表示
	fmt.Printf("\n 🔍️ '%s' の検索結果:\n\n", keyword)
	for i, repo := range result.Items {
		fmt.Printf("%d. %s\n", i+1, repo.Name)
		fmt.Printf("   🌟 Stars: %d\n", repo.Stars)
		fmt.Printf("   🔗 URL:   %s\n", repo.HTMLURL)
		if repo.Description != "" {
			fmt.Printf("   🗒️ %s\n", repo.Description)
		}
		fmt.Println()

		// ログにも書き込む
		log.Printf("取得: %s (%d stars) [%s]", repo.Name, repo.Stars, repo.HTMLURL)
	}

	log.Println("アプリケーション終了:", time.Now().Format(time.RFC3339))
}
