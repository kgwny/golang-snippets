package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/url"
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

func main() {
	http.HandleFunc("/", handleForm)
	http.HandleFunc("/search", handleSearch)

	log.Println("http://localhost:8080 で起動中...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	tmpl := `
		<!DOCTYPE html>
		<html>
		<head><meta charset="utf-8"><title>GitHub検索</title></head>
		<body>
			<h1>GitHubリポジトリ検索</h1>
			<form action="/search" method="get">
				<input type="text" name="q" placeholder="キーワード">
				<button type="submit">検索</button>
			</form>
		</body>
		</html>
	`
	w.Write([]byte(tmpl))
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	escapedQuery := url.QueryEscape(query)
	count := "10"
	apiURL := "https://api.github.com/search/repositories?q=" + escapedQuery + "&sort=stars&order=desc&per_page=" + count

	resp, err := http.Get(apiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "GitHub APIへのアクセスに失敗しました", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var result GitHubResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		http.Error(w, "レスポンスの解析に失敗しました", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		http.Error(w, "テンプレートの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}

	data := struct {
		Query string
		Repos []Repository
	}{
		Query: query,
		Repos: result.Items,
	}

	tmpl.Execute(w, data)
}
