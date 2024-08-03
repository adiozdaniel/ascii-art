package tests

import "github.com/adiozdaniel/ascii-art/pkg/utils"

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
	app = &utils.Input{}
)
