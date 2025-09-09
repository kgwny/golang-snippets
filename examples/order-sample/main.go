package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Order モデル
type Order struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	CustomerName string  `json:"customer_name"`
	Product      string  `json:"product"`
	Quantity     int     `json:"quantity"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
}

// DB グローバル変数
var db *gorm.DB

func main() {
	// DB接続情報
	dsn := "user:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

	// DB接続
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// マイグレーション
	if err := db.AutoMigrate(&Order{}); err != nil {
		log.Fatal("failed to migrate:", err)
	}

	// Echo インスタンス
	e := echo.New()

	// CRUD ルート
	e.POST("/orders", createOrder)
	e.GET("/orders", getOrders)
	e.GET("/orders/:id", getOrder)
	e.PUT("/orders/:id", updateOrder)
	e.DELETE("/orders/:id", deleteOrder)

	// サーバ起動
	e.Logger.Fatal(e.Start(":8080"))
}

// Create
func createOrder(c echo.Context) error {
	order := new(Order)
	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := db.Create(order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, order)
}

// Read (全件)
func getOrders(c echo.Context) error {
	var orders []Order
	if err := db.Find(&orders).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, orders)
}

// Read (1件)
func getOrder(c echo.Context) error {
	id := c.Param("id")
	var order Order
	if err := db.First(&order, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
	}
	return c.JSON(http.StatusOK, order)
}

// Update
func updateOrder(c echo.Context) error {
	id := c.Param("id")
	var order Order
	if err := db.First(&order, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
	}
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := db.Save(&order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, order)
}

// Delete
func deleteOrder(c echo.Context) error {
	id := c.Param("id")
	if err := db.Delete(&Order{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
