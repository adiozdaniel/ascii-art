package main

type args struct {
	input        []string
	fileContents []string
}

type tests struct {
	name     string
	args     args
	expected interface{}
}
