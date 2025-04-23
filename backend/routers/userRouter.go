package routers

import (
	controllers "backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouters(incomingrouters *gin.RouterGroup) {
	users := incomingrouters.Group("/users")
	users.Use(middleware.Authenticate())
	users.GET("/users", controllers.GetUsers())
	users.GET("/users/:user_id", controllers.GetUser())
	users.GET("/leaderboard", controllers.GetRank())
	users.GET("/game/random", controllers.GenerateAnswer())
	users.POST("/game/guess", controllers.Guess())
}
