package responses

import "github.com/gin-gonic/gin"

func ErrorResponse(statusCode int, message interface{}) gin.H {
	return gin.H{
		"status":  statusCode,
		"message": message,
		"success": false,
	}
}
