package config

import (
	"sync"
)

// StateManager manages the application state.
type StateManager struct {
	input  *InputData
	config *AppConfig
	td     *TemplateData
}

// singleton instance of StateManager
var (
	instance *StateManager
	once     sync.Once
)

// GetStateManager returns the singleton instance of StateManager
func GetStateManager() *StateManager {
	once.Do(func() {
		instance = &StateManager{}
		instance.input = NewInputData()
		instance.config = App()
		instance.td = NewTemplateData()
	})
	return instance
}

// GetInput returns the Input instance of StateManager
func (sm *StateManager) GetInput() *InputData {
	return sm.input
}

// GetConfig returns the AppConfig instance of StateManager
func (sm *StateManager) GetConfig() *AppConfig {
	return sm.config
}
