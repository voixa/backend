// File: handlers.go
// Functionalities: Define behavior for routes.

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "I'm alive!"})
}
