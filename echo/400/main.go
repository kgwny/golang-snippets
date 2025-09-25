package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// 400 Bad Request の例:
// サーバー起動後に http://localhost:8080/check-age にアクセスする
func main() {
	e := echo.New()

	// サンプル: 年齢をクエリパラメータで受け取るAPI
	e.GET("/check-age", func(c echo.Context) error {
		ageParam := c.QueryParam("age")

		// age が指定されていない場合
		// http://localhost:8080/check-age にアクセスする
		// {"error":"missing required parameter: age"}
		if ageParam == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "missing required parameter: age",
			})
		}

		// 数値に変換できるかチェック
		// ex) http://localhost:8080/check-age?age=aaa でアクセスする
		// {"error":"invalid parameter: age must be a number"}
		age, err := strconv.Atoi(ageParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid parameter: age must be a number",
			})
		}

		// 正常系
		// ex) http://localhost:8080/check-age?age=100 でアクセスする
		// {"age":100,"status":"valid request"}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"age":    age,
			"status": "valid request",
		})
	})

	// サンプル: 強制的に400を返すエンドポイント
	e.GET("/force400", func(c echo.Context) error {
		return c.NoContent(http.StatusBadRequest)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

// c.JSON(http.StatusBadRequest, {...})
// JSON形式でエラー理由を返す。

// c.NoContent(http.StatusBadRequest)
// レスポンスボディなしで 400 を返す。

// 実務では「エラーコード」「エラーメッセージ」「詳細情報（フィールド名など）」
// を含めた統一フォーマットを定義すると便利です。
