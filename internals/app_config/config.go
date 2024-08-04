package appconfig

import (
	"sync"

	"github.com/adiozdaniel/ascii-art/pkg/utils"
)

// App state
var (
	instance *utils.Input
	once     sync.Once
)

// GetState returns the App state
func GetState() *utils.Input {
	once.Do(func() {
		instance = utils.NewInput()
	})
	return instance
}
