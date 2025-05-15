package main

import (
	"kotoshop/handlers"
	"kotoshop/postgres"
	"log"
	"os"
	"time"

	_ "kotoshop/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Kotoshop
// @version         2.0
// @description     This is the coolest kotoshop
// @host      localhost:8080
// @securityDefinitions.basic  BasicAuth
func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	postgres.Open(os.Getenv("POSTGRES_STRING"))
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/api/auth/signup", handlers.Signup)
	r.POST("/api/auth/login", handlers.Login)
	r.GET("/api/auth/profile", handlers.AuthMiddleware,handlers.Profile)
	r.PUT("/api/auth/update", handlers.AuthMiddleware, handlers.UpdateUser)

	r.POST("/api/products/post", handlers.CreateProduct)
	r.GET("/api/products/get_all", handlers.GetAllProducts)

	r.POST("/api/feedback/post", handlers.AuthMiddleware, handlers.PostFeedback)
	r.GET("/api/feedback/get_all", handlers.GetFeedbacks)
	r.GET("/api/feedback/get_feedback", handlers.AuthMiddleware, handlers.GetUserFeedback)
	r.PUT("/api/feedback/update_feedback", handlers.AuthMiddleware, handlers.UpdateFeedback)

	r.POST("/api/cart/add_product", handlers.AuthMiddleware, handlers.AddToCart)
	r.GET("/api/cart/get_cart", handlers.AuthMiddleware, handlers.GetCart)
	r.PUT("/api/cart/remove_product", handlers.AuthMiddleware, handlers.DeleteCartItem)
	r.DELETE("/api/cart/clean_cart", handlers.AuthMiddleware, handlers.CleanCart)
	
	r.POST("/api/order/create", handlers.AuthMiddleware, handlers.CreateOrder)
	r.GET("/api/order/get_all", handlers.AuthMiddleware, handlers.GetUserOrders)

	r.GET("/api/image/get", handlers.AuthMiddleware, handlers.GetProductImage)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDYzOTE4NzQsImlhdCI6MTc0Mzc5OTg3NCwiaXNzIjoidG9kby1hcHAiLCJzdWIiOjF9.7pDh3AJVygRo4mhSxGY2sDQOfZsdNVnQJyaWeYouRPY