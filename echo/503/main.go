package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 503 Service Temporarily Unavailable の例:
// サーバー起動後に http://localhost:8080/data にアクセスする

// Go + Echo で 503 Service Unavailable を返すのは、典型的には
// ・外部サービスやDBが落ちていて利用できない
// ・サーバーが一時的に過負荷状態
// ・メンテナンス中
// といった場合です。
func main() {
	e := echo.New()

	// サンプル: 外部サービスに依存するAPI
	e.GET("/data", func(c echo.Context) error {
		// ダミー: 外部サービスが落ちていると仮定
		externalServiceAvailable := false

		if !externalServiceAvailable {
			err := errors.New("external service unavailable")

			// 内部ログには詳細を残す
			// {"time":"2025-09-16T10:18:11.369844+09:00",
			// "level":"ERROR",
			// "prefix":"echo",
			// "file":"main.go",
			// "line":"30",
			// "message":"external service unavailable"}
			c.Logger().Error(err)

			// クライアントには503を返却する
			// {"error":"service temporarily unavailable"}
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"error": "service temporarily unavailable",
			})
		}

		// 正常系(サービスが利用可能な場合)
		return c.JSON(http.StatusOK, map[string]string{
			"message": "data response",
		})
	})

	// サンプル: 強制的に503を返却するエンドポイント
	e.GET("/maintenance", func(c echo.Context) error {
		// Retry-After ヘッダをつけると親切（秒数またはHTTP日時形式）
		c.Response().Header().Set("Retry-After", "60") // 60秒後に再試行を推奨
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"error": "server under maintenance, please try again later",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}

// http.StatusServiceUnavailable は 503 を意味する。
// 503 の場合、Retry-After ヘッダを付与するのがベストプラクティス。
// 内部のエラー詳細はログに残し、クライアントにはシンプルなメッセージを返す。
