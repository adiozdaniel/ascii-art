package main

import (
	ascii "ascii/functions"
	"fmt"
	"os"
	"strings"
)

func main(){
	arg := os.Args
	if len(os.Args) > 3 {
		fmt.Println("Incorrect number of arguments")
		return
	}
	if len(os.Args) == 1 {
		fmt.Println("Incorrect number of arguments")
		return
	}
	if arg[1] == "" {
		return
	}
	if arg[1] == "\\n" {
		fmt.Println()
		return
	}

	new := strings.Split(arg[1], "\\n")

	filename := ""


	input , err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	
	var input2 []string
	if filename == "thinkertoy.txt" {
		input2 = strings.Split(string(input), "\r\n")
	} else {
		input2 = strings.Split(string(input), "\n")
	}

	output := ascii.Writer(new, input2)

	fmt.Print(output)
}