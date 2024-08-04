package tests

import (
	appconfig "github.com/adiozdaniel/ascii-art/internals/app_config"
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

// global variables used in the tests
var (
	app = appconfig.GetState()
)
