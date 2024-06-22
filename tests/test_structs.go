package main


// Structs used to store the data in test files
type args struct {
	input        []string
	fileContents []string
	color        string
	reff         string
}

type tests struct {
	name     string
	args     args
	expected interface{}
}