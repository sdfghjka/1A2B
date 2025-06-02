package middleware

import (
	"backend/service"
	// "database/sql"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InjectUserService(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userService := service.NewUserService(db)
		c.Set("userService", userService)
		c.Next()
	}
}
