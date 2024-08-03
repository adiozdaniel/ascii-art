package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Input holds all the data passed around in the application
type Input struct {
	Flags      map[string]string
	BannerFile map[string]string
	ValidFlags map[string]bool
	Args       []string
}

// NewInput creates a new Input instance with default values
func NewInput() *Input {
	return &Input{
		Flags:      members,
		BannerFile: bannerFiles,
		ValidFlags: validFlags,
	}
}

// bannerFiles is a map for banner files and their paths
var bannerFiles = map[string]string{
	"-shadow":      "shadow.txt",
	"--shadow":     "shadow.txt",
	"shadow":       "shadow.txt",
	"-thinkertoy":  "thinkertoy.txt",
	"--thinkertoy": "thinkertoy.txt",
	"thinkertoy":   "thinkertoy.txt",
	"-standard":    "standard.txt",
	"--standard":   "standard.txt",
	"standard":     "standard.txt",
}

// members holds struct members with default values
var members = map[string]string{
	"font":   "--standard",
	"input":  "",
	"color":  "",
	"reff":   "",
	"align":  "left",
	"output": "",
	"isWeb":  "",
}

// validFlags stores allowed flags
var validFlags = map[string]bool{
	"--color":      true,
	"-color":       true,
	"--align":      true,
	"-align":       true,
	"--output":     true,
	"-output":      true,
	"--shadow":     true,
	"-shadow":      true,
	"--thinkertoy": true,
	"-thinkertoy":  true,
	"--standard":   true,
	"-standard":    true,
	"--reff":       true,
	"-reff":        true,
}

// Init initializes the Input
func (i *Input) Init() {
	if len(os.Args) < 1 {
		i.ErrorHandler("fatal")
	}
	*i = *NewInput()
}

// BannerFiles returns the map of banner files for the Input instance
func (i *Input) BannerFiles() map[string]string {
	return i.BannerFile
}

// GetBannerPath returns the name for a specific banner key or defaults to "standard.txt"
func (i *Input) GetBannerPath(key string) string {
	if name, ok := i.BannerFile[key]; ok {
		return name
	}
	return "standard.txt"
}

// Validate checks if the Input contains valid arguments and flags
func (i *Input) Validate() error {
	if i.Flags["output"] != "" && !strings.HasSuffix(i.Flags["output"], ".txt") {
		return fmt.Errorf("output file must have a .txt extension")
	}
	return nil
}

// ParseArgs parses command-line arguments and sets Input fields
func (i *Input) ParseArgs() {
	if i.Flags == nil {
		i.Flags = make(map[string]string)
	}

	for j := 0; j < len(i.Args); {
		input := i.Args[j]
		parts := strings.SplitN(input, "=", 2)
		if len(parts) != 2 {
			j++
			continue
		}

		flag := parts[0]
		value := parts[1]

		if i.IsValidFlag(flag) {
			parsedFlag := i.RemoveLeadingDashes(flag)
			i.Flags[parsedFlag] = value
			i.Args = append(i.Args[:j], i.Args[j+1:]...)
		} else if strings.HasPrefix(flag, "-") {
			i.Args = append(i.Args[:j], i.Args[j+1:]...)
		} else {
			j++
		}
	}

	if err := i.Validate(); err != nil {
		i.ErrorHandler(err.Error())
	}

	i.Flags["input"] = strings.Join(i.Args, " ")
}

// IsValidFlag checks if a flag is valid
func (i *Input) IsValidFlag(flag string) bool {
	return i.ValidFlags[flag]
}

// RemoveLeadingDashes removes leading '--' from the given string
func (i *Input) RemoveLeadingDashes(input string) string {
	if strings.HasPrefix(input, "--") {
		return input[2:]
	}
	return input
}

// RemoveQuotes removes opening or closing quotes in a string
func (i *Input) RemoveQuotes(input string) string {
	var result strings.Builder
	inQuotes := false

	for _, ch := range input {
		if ch == '"' {
			inQuotes = !inQuotes
		} else if ch != '\'' {
			result.WriteRune(ch)
		}
	}

	return strings.TrimSpace(result.String())
}

// GetProjectRoot dynamically finds the project root directory
func (i *Input) GetProjectRoot(path, name string) string {
	cwd, _ := os.Getwd()
	baseDir := cwd
	if strings.HasSuffix(baseDir, "/web") || strings.HasSuffix(baseDir, "/cli") {
		baseDir = filepath.Join(cwd, "../../")
	}
	return filepath.Join(baseDir, path, name)
}
