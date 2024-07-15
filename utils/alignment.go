package utils

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

/*
Alignment formats the given ASCII art on the specified justification and terminal width.
It removes ANSI escape codes from the input before applying alignment.
Supported justifications: "left", "center", "right".
Inputs.Justify should be set to the desired justification mode.
*/
func Alignment(output string, width int) string {
	justification := Inputs.Justify

	if width == 0 {
		width = 80 // fallback to default width
	}

	switch justification {
	case "center":
		return centerAlign(output, width)
	case "right":
		return rightAlign(output, width)
	// case "justify":
	// return justifyAlign(fileContents, ascii_map, output, width)
	default:
		return leftAlign(output)
	}
}

/*
getTerminalWidth retrieves the current width of the terminal.
It uses a system call to obtain terminal size information.
Returns: The width of the terminal in columns. Zero if the width cannot be determined.
*/
func GetTerminalWidth() int {
	type winsize struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}

	ws := &winsize{}
	retCode, _, _ := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		ErrorHandler("align")
	}

	return int(ws.Col)
}

/*
leftAlign aligns output to the left within the terminal width.
Parameters:

	output: output to be left-aligned.

Returns: output as is.
*/
func leftAlign(output string) string {
	return output
}

/*
centerAlign centers output within the terminal width by padding spaces.
Parameters:

	output: output to be centered.
	width: Width of the terminal.

Returns: Centered output.
*/
func centerAlign(output string, width int) string {
	lines := strings.Split(output, "\n")
	var justifiedLines []string
	for _, line := range lines {
		cleanLine := removeANSICodes(line)
		padding := (width - len(cleanLine)) / 2
		if padding < 0 {
			padding = 0
		}
		justifiedLines = append(justifiedLines, fmt.Sprintf("%s%s", strings.Repeat(" ", padding), line))
	}
	return strings.Join(justifiedLines, "\n")
}

/*
rightAlign aligns output to the right within the terminal width by padding spaces.
Parameters:

	output: output to be right-aligned.
	width: Width of the terminal.

Returns: Right-aligned output.
*/
func rightAlign(output string, width int) string {
	lines := strings.Split(output, "\n")
	var justifiedLines []string
	for _, line := range lines {
		cleanLine := removeANSICodes(line)
		padding := width - len(cleanLine)
		if padding < 0 {
			padding = 0
		}
		justifiedLines = append(justifiedLines, fmt.Sprintf("%s%s", strings.Repeat(" ", padding), line))
	}
	return strings.Join(justifiedLines, "\n")
}

/*
justifyAlign centers output within the terminal width by inserting spaces between words.
Parameters:

	output: output to be justified.
	width: Width of the terminal.

Returns: Justified output.
*/
func justifyAlign(fileContents []string, ascii_map map[rune]int, output string, width int) string {
	var justifiedLine strings.Builder
	var words = strings.Fields(removeANSICodes(output))

	wordsLength := 0
	for _, word := range words {
		wordsLength += len(word)
	}

	totalSpaces := width - wordsLength
	spaceSlots := len(words) - 1
	evenSpaces := totalSpaces / spaceSlots
	extraSpaces := totalSpaces % spaceSlots

	for _, line := range strings.Split(Inputs.Input, "\n") {
		for i := 0; i < 8; i++ {
			var builder strings.Builder
			for _, char := range line {
				if ascii, ok := ascii_map[char]; ok {
					builder.WriteString(fileContents[ascii+i])
				}
				if i < spaceSlots {
					builder.WriteString(strings.Repeat(" ", evenSpaces))
					if i < extraSpaces {
						builder.WriteString(" ")
						extraSpaces--
					}
				}
			}
			justifiedLine.WriteString(builder.String())
			justifiedLine.WriteRune('\n')
		}
		justifiedLine.WriteRune('\n')
	}

	return justifiedLine.String()
}
