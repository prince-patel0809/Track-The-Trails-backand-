package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDashboardHandler(c *gin.Context) {

	userID := c.MustGet("userID").(string)

	data, err := GetDashboard(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch dashboard",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"dashboard": data,
	})
}
