package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

// 環境変数
type appEnvironment struct {
	username string
	password string
}

// レスポンスJsonの定義
type apiRequestResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// DBのフィールド
type person struct {
	id   int
	name string
}

func main() {
	e := echo.New()
	e.GET("/request", makeHandler(requestHandler))
	e.GET("/person", makeHandler(personHandler))
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Logger.Fatal(e.Start(":8080"))
}

// handlerに共通の変数を渡すため、2つの引数をもつfunctionを引数にとり、func(c echo.Context)を返す
func makeHandler(fn func(c echo.Context, env appEnvironment) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return fn(c, getAppEnvironment())
	}
}

// APIを叩いてレスポンスを返す
func requestHandler(c echo.Context, env appEnvironment) error {
	// クエリパラメータ取得＆バリデーションチェック
	paramKey := c.QueryParam("key")
	if len(paramKey) == 0 {
		return c.JSON(http.StatusOK, apiRequestResponse{Code: 2002, Message: "key parameter not found."})
	}

	// HTTPリクエストを行い結果をログに出力
	resp, _ := http.Get("http://httpbin.org/get?key=" + paramKey)
	if resp == nil {
		return c.JSON(http.StatusOK, apiRequestResponse{Code: 2003, Message: "HTTP Request failed."})
	}
	defer resp.Body.Close()
	byteArray, _ := io.ReadAll(resp.Body)
	fmt.Printf("log: %s", byteArray)

	return c.JSON(http.StatusOK, apiRequestResponse{Code: 2000, Message: "OK"})
}

// DBを叩いてレスポンスを返す
func personHandler(c echo.Context, env appEnvironment) error {
	// クエリパラメータ取得、バリデーションチェック
	paramKey := c.QueryParam("user")
	if len(paramKey) == 0 {
		return c.JSON(http.StatusOK, apiRequestResponse{Code: 2002, Message: "user parameter not found."})
	}

	// DBからデータを取得して結果をログに出力
	db, err := sql.Open("mysql", "root:"+env.password+"@/"+env.username)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var person person
		err := rows.Scan(&person.id, &person.name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(person.id, person.name)
	}
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
	return c.JSON(http.StatusOK, apiRequestResponse{Code: 2000, Message: "OK"})
}

// エラーレスポンスのカスタマイズ
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if code == http.StatusMethodNotAllowed {
		// 405エラーの場合にレスポンスをカスタマイズする
		c.JSON(code, apiRequestResponse{Code: 2001, Message: "Method not allowed."})
	} else {
		// 405以外のエラーが発生した場合は以下の内容でレスポンスする
		c.JSON(code, apiRequestResponse{Code: 5000, Message: "Internal server error."})
	}
}

// 環境変数を取得
func getAppEnvironment() appEnvironment {
	if len(os.Getenv("DB_USERNAME")) == 0 || len(os.Getenv("DB_PASSWORD")) == 1 {
		panic("appEnvironment not setting.")
	}

	return appEnvironment{
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
	}
}
