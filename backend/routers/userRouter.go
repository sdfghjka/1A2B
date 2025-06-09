package routers

import (
	controllers "backend/controllers"

	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouters(incomingrouters *gin.RouterGroup) {
	incomingrouters.POST("/ai/start", controllers.AIVersionConnect)
	users := incomingrouters.Group("/users")
	users.Use(middleware.Authenticate())
	users.GET("/admin", controllers.GetUsers())
	users.GET("/:user_id", controllers.GetUser())
	users.GET("/me", controllers.GetInfo())
	users.GET("/leaderboard", controllers.GetRank())
	users.GET("/game/random", controllers.GenerateAnswer())
	users.POST("/game/guess", controllers.Guess())
	users.POST("/upload/image", controllers.UploadUserImage)
}
