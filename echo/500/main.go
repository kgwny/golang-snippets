package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 500 Internal Server Error の例:
// サーバー起動後に http://localhost:8080/user/999 にアクセスする
// {"error":"internal server error"}
// Go + Echo で 500 Internal Server Error を返すのは、
// ・DB クエリ失敗
// ・外部 API 呼び出しで異常が発生
// ・サーバー内部の予期せぬエラー
// といったケースです。
func main() {
	e := echo.New()

	// サンプル: ユーザー情報を取得するエンドポイント
	e.GET("/user/:id", func(c echo.Context) error {
		id := c.Param("id")

		// ダミー: id が "999" のときにエラーが発生する前提としたとき
		if id == "999" {
			// ここでDBエラーが起きたと仮定する
			err := errors.New("database connection failed")

			// ログにエラーを出力する
			c.Logger().Error(err)

			// クライアントには 500 を返却する
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "internal server error",
			})
		}

		// 正常ケース: サーバー起動後に http://localhost:8080/user/100 にアクセスする
		// {"id":"100","name":"Taro"}
		return c.JSON(http.StatusOK, map[string]string{
			"id":   id,
			"name": "Taro",
		})
	})

	// サンプル: 強制的に500を返却するエンドポイント
	e.GET("/force500", func(c echo.Context) error {
		return c.NoContent(http.StatusInternalServerError)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

// c.JSON(http.StatusInternalServerError, {...})
// エラーメッセージを JSON で返す。

// c.NoContent(http.StatusInternalServerError)
// ボディなしで返す。

// 実運用では、内部エラーの詳細はログにのみ出力し、
// クライアントには一般的なメッセージ（"internal server error"）だけ返すのがベストプラクティス。
