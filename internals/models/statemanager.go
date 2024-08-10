package models

import (
	"net/url"
	"sync"
)

// StateManager manages the application state.
type StateManager struct {
	input   *InputData
	config  *AppConfig
	td      *TemplateData
	session *SessionManager
	forms   *Forms
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
		instance.session = NewSessionManager()
		instance.forms = NewForms(make(url.Values))
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

// GetTemplateData returns the TemplateData instance of StateManager
func (sm *StateManager) GetTemplateData() *TemplateData {
	return sm.td
}

// GetSessionManager returns the SessionManager instance of StateManager
func (sm *StateManager) GetSessionManager() *SessionManager {
	return sm.session
}

// GetForms returns the Forms instance of StateManager
func (sm *StateManager) GetForms() *Forms {
	return sm.forms
}
