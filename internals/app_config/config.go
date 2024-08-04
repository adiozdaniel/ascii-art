package appconfig

import (
	"sync"

	"github.com/adiozdaniel/ascii-art/pkg/utils"
)

// StateManager manages the application state.
type StateManager struct {
	input *utils.Input
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
		instance.input = utils.NewInput()
	})
	return instance
}

// GetInput returns the Input instance of StateManager
func (sm *StateManager) GetInput() *utils.Input {
	return sm.input
}
