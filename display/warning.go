package display

import "fmt"

func WarningPrompt(r string)  {
	fmt.Println(Warning)
	fmt.Printf("The character '%s' is not part of our ASCII graphics. It has been skipped", r)
}