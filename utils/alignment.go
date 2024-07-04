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
func Alignment(asciiArt string) {
	justification := Inputs.Justify
	width := getTerminalWidth()

	if width == 0 {
		width = 80 // fallback to default width
	}

	switch justification {
	case "center":
		fmt.Println(centerJustify(asciiArt, width))
	case "right":
		fmt.Println(rightJustify(asciiArt, width))
	default:
		fmt.Println(leftJustify(asciiArt))
	}
}

/*
getTerminalWidth retrieves the current width of the terminal.
It uses a system call to obtain terminal size information.
Returns: The width of the terminal in columns. Zero if the width cannot be determined.
*/
func getTerminalWidth() int {
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
leftJustify aligns text to the left within the terminal width.
Parameters:

	text: Text to be left-aligned.

Returns: text as is.
*/
func leftJustify(text string) string {
	return text
}

/*
centerJustify centers text within the terminal width by padding spaces.
Parameters:

	text: Text to be centered.
	width: Width of the terminal.

Returns: Centered text.
*/
func centerJustify(text string, width int) string {
	lines := strings.Split(text, "\n")
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
rightJustify aligns text to the right within the terminal width by padding spaces.
Parameters:

	text: Text to be right-aligned.
	width: Width of the terminal.

Returns: Right-aligned text.
*/
func rightJustify(text string, width int) string {
	lines := strings.Split(text, "\n")
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
