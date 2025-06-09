package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func RespondError(c *gin.Context, err any, code int) {
	c.AbortWithStatusJSON(code, gin.H{"error": err})
}
