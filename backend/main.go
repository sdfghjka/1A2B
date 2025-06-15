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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://e312-2401-e180-8820-8783-ed6c-8b9f-c725-1ce2.ngrok-free.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	db := database.MysqlDB()
	GameService := service.NewGameService(db)
	defer db.Close()
	database.InitRedis()
	service.InitOAuthProviders()
	api := router.Group("/api")
	api.GET("/ws", controllers.JoinRoomHandler(GameService))
	api.GET("/ai/start", controllers.StartAIGameHandler(GameService))
	router.Use(gin.Logger(), middleware.ErrorHandler(), middleware.InjectUserService(db))
	api.POST("/payment", controllers.Payment)
	routers.AuthRouters(api)
	routers.UserRouters(api)
	router.Run(":" + port)
}
