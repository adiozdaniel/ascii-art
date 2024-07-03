package utils

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

func Display(asciiArt string) {
	// 	asciiArt := `
	//     _.-;;-._
	//  '-..-'|   ||   |
	//  '-..-'|_.-;;-._|
	//  '-..-'|   ||   |
	//  '-..-'|_.-''-._|`

	// Get the terminal width dynamically
	width := getTerminalWidth()
	justification := "center" // change to "left", "center", or "right" as needed

	// Justify and print ASCII art
	switch justification {
	case "center":
		fmt.Println(centerJustify(asciiArt, width))
	case "right":
		fmt.Println(rightJustify(asciiArt, width))
	default:
		fmt.Println(leftJustify(asciiArt, width))
	}
}

func getTerminalWidth() int {
	type winsize struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}

	width := 0
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		fmt.Println("Error getting terminal size:", errno)
		return 0
	}

	width = int(ws.Col)

	if width == 0 {
		width = 80 // fallback to default width
	}
	return width
}

func leftJustify(text string, width int) string {
	lines := strings.Split(text, "\n")
	var justifiedLines []string
	for _, line := range lines {
		justifiedLines = append(justifiedLines, line)
	}
	return strings.Join(justifiedLines, "\n")
}

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
