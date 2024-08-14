package helpers

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
	alignment := app.Flags["align"]

	if !isValidAlignment(alignment) {
		return ""
	}

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
	case "left":
		return leftAlign(output, width)
	default:
		app.ErrorHandler("justify")
	}
	return ""
}

/*
getTerminalWidth retrieves the current width of the terminal.
It uses a system call to obtain terminal size information.
Returns: The width of the terminal in columns. Zero if the width cannot be determined.
*/
func GetTerminalWidth() (int, int) {
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
		app.ErrorHandler("align")
	}

	return int(ws.Row), int(ws.Col)
}

/*
leftAlign aligns output to the left within the terminal width.
Parameters:

	output: output to be left-aligned.

Returns: output as is.
*/
func leftAlign(output string, width int) string {
	lines := strings.Split(output, "\n")
	var leftLines []string
	for _, line := range lines {
		cleanLine := RemoveANSICodes(line)
		if len(cleanLine) <= width {
			leftLines = append(leftLines, strings.ReplaceAll(string(line), "$", "  "))
		}
	}
	return strings.Join(leftLines, "\n")
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
		cleanLine := RemoveANSICodes(line)
		padding := (width - len(cleanLine)) / 2
		if padding < 0 {
			padding = 0
		}

		if len(cleanLine) <= width {
			centeredLines = append(centeredLines, fmt.Sprintf("%s%s", strings.Repeat(" ", padding), strings.ReplaceAll(string(line), "$", "  ")))
		}
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
		cleanLine := RemoveANSICodes(line)
		padding := width - len(cleanLine) - 1
		if padding < 0 {
			padding = 0
		}

		if len(cleanLine) <= width {
			rightLines = append(rightLines, fmt.Sprintf("%s%s", strings.Repeat(" ", padding), strings.ReplaceAll(line, "$", "  ")))
		}
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
		cleanLine := RemoveANSICodes(line)
		slots, ln := spaceSlots(cleanLine)
		givenSpaces := width - ln
		var spacePerSlot int

		if slots == 0 {
			justifiedLines = append(justifiedLines, line)
		} else if len(cleanLine) <= width {
			spacePerSlot = givenSpaces / slots
			justifiedLines = append(justifiedLines, strings.ReplaceAll(string(line), "$", strings.Repeat(" ", spacePerSlot)))
		}
	}

	return strings.Join(justifiedLines, "\n")
}

/*
spaceSlots calculates the number of spaces and words in the given output.
Parameters:

	output: Output to be analyzed.

Returns: Number of spaces and length of the output.
*/
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
