package responses

import (
	"github.com/gin-gonic/gin"
)

func RespondData(c *gin.Context, data any, code int) {
	c.JSON(code, gin.H{"data": data})
}

func RespondError(c *gin.Context, err any, code int) {
	c.AbortWithStatusJSON(code, gin.H{"error": err})
}
