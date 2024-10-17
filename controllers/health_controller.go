package controllers

import (
	"github.com/BeloTechnologies/session-core/core_models"
	"github.com/gin-gonic/gin"
)

// CheckHealth will return a success response if the server is running
func CheckHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, core_models.SuccessResponse{
			Message: "Service is up and running",
			Status:  200,
		})
	}
}
