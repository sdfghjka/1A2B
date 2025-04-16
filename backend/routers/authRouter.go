package routers

import (
	controllers "backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouters(incomingrouters *gin.RouterGroup) {
	user := incomingrouters.Group("/user")
	user.POST("/signup", controllers.Signup())
	user.POST("/login", controllers.Login())
}
