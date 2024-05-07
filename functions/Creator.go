package ascii

import "fmt"

func Writer(new, input2 []string) string {
	var height int
	var line string

	for _, word := range new {
		if word == "" {
			height = 1
		} else {
			height = 8
		}
		for i := 1 ; i <= height; i++ {
			for _, ch := range word {
				if ch < 32 || ch > 126 {
					fmt.Println("Error: Input contains non-ascii cahrachters")
					return ""
				}
				index := (ch-32)*9 + rune(i)
				line += input2[index]
			}
			line += "\n"
		}
	}
	return line
}