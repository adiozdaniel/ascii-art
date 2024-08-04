package tests

import (
	"github.com/adiozdaniel/ascii-art/internals/models"
)

// Structs used to store the data in test files
type args struct {
	input        string
	fileContents []string
}

type tests struct {
	name     string
	args     args
	expected interface{}
}

// get the app state manager
var (
	sm  = models.GetStateManager()
	app = sm.GetInput()
)
