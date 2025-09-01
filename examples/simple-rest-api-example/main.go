package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// モデル定義
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"uniqueIndex"`
}

// グローバル変数
var db *gorm.DB

// DB初期化
func initDB() {
	dsn := "user:password@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// マイグレーション実行
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal("failed to migrate:", err)
	}
}

// ハンドラー

// Create User
func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if result := db.Create(&user); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

// Get All Users
func getUsers(c echo.Context) error {
	var users []User
	if result := db.Find(&users); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

// Get User by ID
func getUser(c echo.Context) error {
	id := c.Param("id")
	var user User
	if result := db.First(&user, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, user)
}

// Update User
func updateUser(c echo.Context) error {
	id := c.Param("id")
	var user User
	if result := db.First(&user, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	updated := new(User)
	if err := c.Bind(updated); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	user.Name = updated.Name
	user.Email = updated.Email
	if result := db.Save(&user); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

// Delete User
func deleteUser(c echo.Context) error {
	id := c.Param("id")
	if result := db.Delete(&User{}, id); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// メイン処理
func main() {
	// DB 初期化
	initDB()

	// Echo インスタンス作成
	e := echo.New()

	// ルーティング
	e.POST("/users", createUser)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// サーバー起動
	port := 8080
	fmt.Printf("Server started at http://localhost:%d\n", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
