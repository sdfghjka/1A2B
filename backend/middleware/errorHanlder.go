package middleware

import (
	httpError "backend/Error"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				apiErr := httpError.FromError(err)
				c.AbortWithStatusJSON(apiErr.StatusCode, gin.H{"error": apiErr.Message})
			}
		}()
		c.Next()
	}
}
