package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	r.GET("/menu", func(ctx *gin.Context) {
		menu := []map[string]any{
			{"id": 1, "name": "Espresso", "price": 150},
			{"id": 2, "name": "Cappuccino", "price": 200},
			{"id": 3, "name": "Latte", "price": 220},
		}
		ctx.JSON(http.StatusOK, gin.H{"menu": menu})
	})

	r.GET("/cart", func(ctx *gin.Context) {
		cart := []map[string]any{
			{"product_id": 2, "name": "Cappuccino", "qty": 1, "price": 200},
			{"product_id": 3, "name": "Latte", "qty": 2, "price": 220},
		}
		ctx.JSON(http.StatusOK, gin.H{"cart": cart, "total": 640})
	})

	fmt.Println("Coffee Shop API is running on :8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Server error:", err)
	}
}
