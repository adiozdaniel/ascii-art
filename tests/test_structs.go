package tests

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
