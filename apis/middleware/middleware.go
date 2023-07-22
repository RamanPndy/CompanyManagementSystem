package middleware

import (
	"fmt"

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
