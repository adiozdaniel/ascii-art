package utilities

import (
	"fmt"
	"os"
)

// ErrorHandler outputs errors and safely exits the program
func ErrorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
