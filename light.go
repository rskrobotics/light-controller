package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LightController interface {
	TurnOn() error
	TurnOff() error
}

func NewLightController() LightController {
	if isMockEnvironment() {
		return &MockLightController{}
	}
	return &RealLightController{}
}

func isMockEnvironment() bool {
	return true
}

type MockLightController struct{}

func (m *MockLightController) TurnOn() error {
	fmt.Println("Mock light turned on")
	return nil
}

func (m *MockLightController) TurnOff() error {
	fmt.Println("Mock light turned off")
	return nil
}

type RealLightController struct{}

func (r *RealLightController) TurnOn() error {
	fmt.Println("Real light turned on")
	return nil
}

func (r *RealLightController) TurnOff() error {
	fmt.Println("Real light turned off")
	return nil
}

func HandleLightState(controller LightController) gin.HandlerFunc {
	return func(c *gin.Context) {
		state := c.Param("state")

		var err error
		switch state {
		case "on":
			err = controller.TurnOn()
		case "off":
			err = controller.TurnOff()
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change light state"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Light state changed successfully"})
	}
}
