package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// 304 Not Modified の例:
// 例えば、If-None-Match や If-Modified-Since といった 条件付きリクエスト を受けたときに、
// 対象リソースが変更されていなければ 304 を返すケースを想定したサンプル
// サーバー起動後に http://localhost:8080/article/100 にアクセスする
func main() {
	e := echo.New()

	// サンプル: 記事取得API
	e.GET("/article/:id", func(c echo.Context) error {
		articleID := c.Param("id")

		// ダミーデータ: 記事の最終更新時刻
		lastModified := time.Date(2025, 9, 1, 12, 0, 0, 0, time.UTC)

		// リクエストヘッダー　If-Modified-Since をチェックする
		ifModifiedSince := c.Request().Header.Get("If-Modified-Since")
		if ifModifiedSince != "" {
			if t, err := http.ParseTime(ifModifiedSince); err == nil {
				// 変更がなければ 304 を返却する
				if !lastModified.After(t) {
					return c.NoContent(http.StatusNotModified)
				}
			}
		}

		// 記事が変更されている場合は本文を返す
		c.Response().Header().Set("Last-Modified", lastModified.UTC().Format(http.TimeFormat))
		return c.JSON(http.StatusOK, map[string]string{
			"id":      articleID,
			"title":   "Echo Framework Guide",
			"content": "This is a sample article.",
		})
	})

	// {"content":"This is a sample article.","id":"100","title":"Echo Framework Guide"}
	e.Logger.Fatal(e.Start(":8080"))
}

// c.NoContent(http.StatusNotModified)
// レスポンスボディなしで 304 を返す。

// 条件付きリクエスト (If-Modified-Since, If-None-Match) をチェックすることで、
// キャッシュが有効な場合に 304 を返せる。

// 304 の場合は レスポンスボディを含めない のが仕様。
