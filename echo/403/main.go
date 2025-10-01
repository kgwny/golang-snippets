package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 403 Forbidden の例:
// サーバー起動後に http://localhost:8080/admin にアクセスする
// {"error":"forbidden: you don't have access to this resource"}
func main() {
	e := echo.New()

	// サンプル: ユーザーが admin でないとアクセスできないエンドポイント
	e.GET("/admin", func(c echo.Context) error {
		// ダミー: ユーザーのロールをリクエストヘッダから取得
		role := c.Request().Header.Get("X-User-Role")

		if role != "admin" {
			// 権限がなければ 403 Forbidden
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "forbidden: you don't have access to this resource",
			})
		}

		// 正常系
		return c.JSON(http.StatusOK, map[string]string{
			"massage": "Welcome, admin!",
		})

	})

	// サンプル: 強制的に 403 を返却するエンドポイント
	e.GET("/force403", func(c echo.Context) error {
		return c.NoContent(http.StatusForbidden)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

// 403 Forbidden を返すケースは、
// 認証は済んでいるが アクセス権限がないとき
// 例えば「一般ユーザーが管理者専用のリソースにアクセスした場合」
// といった場面です。

// 401 Unauthorized は「認証に失敗」したとき。
// 403 Forbidden は「認証済みだが権限が不足」しているとき。

// c.JSON(http.StatusForbidden, {...})
// JSON形式で理由を返す。

// c.NoContent(http.StatusForbidden)
// ボディなしで返す。
