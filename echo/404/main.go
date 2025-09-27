package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 404 Not Found の例:
// サーバー起動後に http://localhost:8080/users/100 にアクセスする
func main() {
	e := echo.New()

	// サンプル: ユーザー取得API
	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")

		// DB検索や処理結果が存在しない場合に404を返却する
		if id != "123" {
			// ダミー条件: idが123以外は存在しない前提とする
			// JSONレスポンスで404
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "User not found",
			})
		}

		// 正常系
		return c.JSON(http.StatusOK, map[string]string{
			"id":   id,
			"name": "Taro",
		})
	})

	// サンプル: 意図的に404を返すエンドポイント
	e.GET("/force404", func(c echo.Context) error {
		// レスポンスボディなしで404
		return c.NoContent(http.StatusNotFound)
	})

	// "message":"Not Found"}
	e.Logger.Fatal(e.Start(":8080"))
}

// c.JSON(http.StatusNotFound, ...)
// JSON でエラーメッセージを返す場合。

// c.NoContent(http.StatusNotFound)
// レスポンスボディなしで 404 を返す場合。

// http.StatusNotFound は定数で 404 と同義。
