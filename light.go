package main

import "fmt"

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

func (m *RealLightController) TurnOn() error {
	fmt.Println("Mock light turned on")
	return nil
}

func (m *RealLightController) TurnOff() error {
	fmt.Println("Mock light turned off")
	return nil
}
