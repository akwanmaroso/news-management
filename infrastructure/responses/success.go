package responses

import "github.com/gin-gonic/gin"

func SuccessResponse(statusCode int, data interface{}) gin.H {
	return gin.H{
		"status":  statusCode,
		"data":    data,
		"success": true,
	}
}
