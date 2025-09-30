package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 401 Unauthorized の例:
// サーバー起動後に http://localhost:8080/secure にアクセスする
func main() {
	e := echo.New()

	// サンプル: 認証が必要なエンドポイント
	e.GET("/secure", func(c echo.Context) error {
		// Authorization ヘッダをチェックする
		token := c.Request().Header.Get("Authorization")

		// トークンがない場合は401を返却する
		// {"error":"missing Authorization header"}
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing Authorization header",
			})
		}

		// トークンが不正な場合は401を返却する
		if token != "Bearer secret-token" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid token",
			})
		}

		// 正常系
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Access granted",
		})
	})

	// サンプル: 強制的に401を返却するエンドポイント
	e.GET("/force401", func(c echo.Context) error {
		return c.NoContent(http.StatusUnauthorized)

	})

	e.Logger.Fatal(e.Start(":8080"))
}

// http.StatusUnauthorized は 401 を意味する。
// 実務では、401 を返すときに WWW-Authenticate ヘッダ を付与することが推奨される（RFC準拠）。
// c.Response().Header().Set("WWW-Authenticate", `Bearer realm="example"`)
// return c.JSON(http.StatusUnauthorized, map[string]string{
//     "error": "unauthorized",
// })
