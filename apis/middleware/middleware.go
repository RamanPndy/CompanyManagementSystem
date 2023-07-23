package middleware

import (
	"fmt"
	"net/http"

	"companybuilder/modules/auth"

	"github.com/gin-gonic/gin"
	pkgErrors "github.com/pkg/errors"
)

// HandlePanic ... rest panic handler
func HandlePanic(c *gin.Context) {
	defer func(c *gin.Context) {
		r := recover()
		var stackTrace string
		if r != nil {
			err, ok := r.(error)
			if ok {
				// Logs the error
				stackTrace = fmt.Sprintf("%+v", pkgErrors.New(err.Error()))

				// Forms error message
				c.JSON(500, gin.H{
					"message":    "Panic: Unexpected error occured.",
					"error":      err.Error(),
					"stackTrace": stackTrace,
				})
			} else {
				// Forms error message
				c.JSON(500, gin.H{
					"message": "Panic: Unexpected error occured, failed to parse error",
				})
			}
		}
	}(c)
	c.Next()
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := auth.TokenValid(ctx.Query("token"), ctx.Request.Header.Get("Authorization"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "request is unauthorized", "description": "authorization token is missing or expired"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
