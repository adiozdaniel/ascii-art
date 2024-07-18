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
	alignment := Inputs.Justify

	if width == 0 {
		width = 80 // fallback to default width
	}

	switch alignment {
	case "center":
		return centerAlign(output, width)
	case "right":
		return rightAlign(output, width)
	case "justify":
		return justifyAlign(output, width)
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
		Row uint16
		Col uint16
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
	var centeredLines []string
	for _, line := range lines {
		cleanLine := removeANSICodes(line)
		padding := (width - len(cleanLine)) / 2
		if padding < 0 {
			padding = 0
		}
		centeredLines = append(centeredLines, fmt.Sprintf("%s%s", strings.Repeat(" ", padding), line))
	}
	return strings.Join(centeredLines, "\n")
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
	var rightLines []string
	for _, line := range lines {
		cleanLine := removeANSICodes(line)
		padding := width - len(cleanLine)
		if padding < 0 {
			padding = 0
		}
		rightLines = append(rightLines, fmt.Sprintf("%s%s", strings.Repeat(" ", padding), line))
	}
	return strings.Join(rightLines, "\n")
}

/*
justifyAlign centers output within the terminal width by inserting spaces between words.
Parameters:

	output: output to be justified.
	width: Width of the terminal.

Returns: Justified output.
*/
func justifyAlign(output string, width int) string {
	lines := strings.Split(output, "\n")
	var justifiedLines []string

	for _, line := range lines {
		var spaceSlots, len = spaceSlots(line)
		var givenSpaces = width - len

		if spaceSlots < 2 {
			justifiedLines = append(justifiedLines, strings.ReplaceAll(line, "$", " "))
			spaceSlots = 1
		}

		var spacePerSlot = givenSpaces / spaceSlots

		if spacePerSlot > 1 {
			justifiedLines = append(justifiedLines, strings.ReplaceAll(string(line), "$", strings.Repeat(" ", spacePerSlot)))
		}
	}

	return strings.Join(justifiedLines, "\n")
}

func spaceSlots(output string) (int, int) {
	var slots int
	var len int

	for i, char := range strings.Split(output, "\n")[0] {
		if char == '$' {
			slots++
		}

		len = i + 1
	}
	return slots, len
}
