package appconfig

import (
	"sync"

	"github.com/adiozdaniel/ascii-art/pkg/utils"
)

// StateManager manages the application state.
type StateManager struct {
	input *utils.Input
	once  sync.Once
}

// GetState returns the singleton instance of StateManager
func GetStateManager() *StateManager {
	var sm StateManager

	sm.once.Do(func() {
		sm.input = utils.NewInput()
	})
	return &sm
}
