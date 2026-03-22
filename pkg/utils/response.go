package utils
import "github.com/gin-gonic/gin"

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}

func Error(c *gin.Context, msg string) {
	c.JSON(500, gin.H{
		"success": false,
		"error":   msg,
	})
}