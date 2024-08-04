package tests

import (
	appconfig "github.com/adiozdaniel/ascii-art/internals/config"
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
	sm  = appconfig.GetStateManager()
	app = sm.GetInput()
)
