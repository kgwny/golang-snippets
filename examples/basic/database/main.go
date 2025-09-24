package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// DSN: "user:password@tcp(host:port)/dbname"
	dsn := "user:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

	// DB接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB接続エラー:", err)
	}
	defer db.Close()

	// 接続確認
	if err := db.Ping(); err != nil {
		log.Fatal("DB疎通確認エラー:", err)
	}
	fmt.Println("DB接続成功")

	// テーブル作成
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50),
		age INT
	);`
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal("テーブル作成エラー:", err)
	}
	fmt.Println("テーブル作成完了")

	// データ登録
	insertStmt, err := db.Prepare("INSERT INTO users(name, age) VALUES(?, ?)")
	if err != nil {
		log.Fatal("INSERT準備エラー:", err)
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec("Alice", 25)
	if err != nil {
		log.Fatal("INSERTエラー:", err)
	}
	_, err = insertStmt.Exec("Bob", 30)
	if err != nil {
		log.Fatal("INSERTエラー:", err)
	}
	fmt.Println("データ挿入完了")

	// データ取得
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal("SELECTエラー:", err)
	}
	defer rows.Close()

	fmt.Println("ユーザー一覧:")
	for rows.Next() {
		var id int
		var name string
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal("Scanエラー:", err)
		}
		fmt.Printf("ID=%d, Name=%s, Age=%d\n", id, name, age)
	}

	// エラーチェック（rows.Next の後）
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
