package main

import (
	"backend/controllers"
	"backend/database"
	"backend/middleware"
	routers "backend/routers"
	"backend/service"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	db := database.MysqlDB()
	defer db.Close()
	database.InitRedis()
	service.InitOAuthProviders()
	router.GET("/api/ws", controllers.JoinRoomHandler)
	router.Use(gin.Logger(), middleware.ErrorHandler(), middleware.InjectUserService(db))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api := router.Group("/api")
	routers.AuthRouters(api)
	routers.UserRouters(api)
	router.Run(":" + port)
}
