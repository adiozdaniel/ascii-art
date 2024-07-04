package utils

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

// Alignment sets the output to the desired alignment
func Alignment(asciiArt string) {
	justification := Inputs.Justify
	width := getTerminalWidth()

	if width == 0 {
		width = 80 // fallback to default width
	}

	switch justification {
	case "center":
		fmt.Println(centerJustify(removeANSICodes(asciiArt), width))
	case "right":
		fmt.Println(rightJustify(removeANSICodes(asciiArt), width))
	default:
		fmt.Println(leftJustify(removeANSICodes(asciiArt)))
	}
}

// getTerminalWidth gets the current terminal width
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

// leftJustify aligns characters to the left of the screen
func leftJustify(text string) string {
	return text
}

// centerJustify aligns characters to the center of the screen
func centerJustify(text string, width int) string {
	lines := strings.Split(text, "\n")
	var justifiedLines []string
	for _, line := range lines {
		padding := (width - len(line)) / 2
		if padding < 0 {
			padding = 0
		}
		justifiedLines = append(justifiedLines, fmt.Sprintf("%s%s", strings.Repeat(" ", padding), line))
	}
	return strings.Join(justifiedLines, "\n")
}

// rightJustify aligns characters to the right of the screen
func rightJustify(text string, width int) string {
	lines := strings.Split(text, "\n")
	var justifiedLines []string
	for _, line := range lines {
		padding := width - len(line)
		if padding < 0 {
			padding = 0
		}
		justifiedLines = append(justifiedLines, fmt.Sprintf("%s%s", strings.Repeat(" ", padding), line))
	}
	return strings.Join(justifiedLines, "\n")
}
