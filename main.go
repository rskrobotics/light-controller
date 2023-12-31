package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	lightController := NewLightController()

	// Route for turning on/off lights
	r.GET("/light/:state/", func(c *gin.Context) {
		state := c.Param("state")

		var err error
		switch state {
		case "on":
			err = lightController.TurnOn()
		case "off":
			err = lightController.TurnOff()
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change light state"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Light state changed successfully."})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
