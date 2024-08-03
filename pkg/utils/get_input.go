package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Input holds all the data passed around in the application
type Input struct {
	Flags      map[string]string
	BannerFile map[string]string
	ValidFlags map[string]bool
	Args       []string
}

// Inputs is a placeholder for the Input struct
var Inputs Input

// NewInput creates a new Input instance with default BannerFiles
func NewInput() *Input {
	return &Input{
		BannerFile: bannerFiles,
		ValidFlags: validFlags,
	}
}

// BannerFiles is a map for banner files and their paths
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

// members holds struct members
var members = map[string]string{
	"color":  "",
	"reff":   "",
	"font":   "",
	"align":  "left",
	"output": "",
	"isWeb":  "",
}

// validFlags stores allowed flags
var validFlags = map[string]bool{
	"--color":      true,
	"--align":      true,
	"--output":     true,
	"-shadow":      true,
	"--shadow":     true,
	"-thinkertoy":  true,
	"--thinkertoy": true,
	"-standard":    true,
	"--standard":   true,
	"--reff":       true,
	"-reff":        true,
}

// init initializes the Input
func (i *Input) Init() {
	if len(os.Args) > 1 {
		i.ErrorHandler("fatal")
	}

	i.Font = "--standard"
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

// RemoveQuotes removes opening or closing quotes in a string
func (i *Input) RemoveQuotes(input string) string {
	var result strings.Builder
	newInput := strings.Fields(input)

	for _, word := range newInput {
		var temp strings.Builder
		var skipNext bool
		var isSpace bool

		for i := 0; i < len(word); i++ {
			if skipNext {
				skipNext = false
				continue
			}

			if word[i] == '=' && i+2 < len(word) && (word[i+1] == '"' || word[i+1] == '\'') {
				temp.WriteByte('=')
				isSpace = true
				skipNext = true
			} else if word[i] == '\\' && i+1 < len(word) && (word[i+1] == '"' || word[i+1] == '\'') {
				if word[i+1] == '"' {
					temp.WriteByte('"')
				}
				if word[i+1] == '\'' {
					temp.WriteByte('\'')
				}
				skipNext = true
			} else if word[i] == '"' || word[i] == '\'' {
				if isSpace {
					temp.WriteByte(word[i])
					isSpace = false
				}
			} else {
				temp.WriteByte(word[i])
			}
		}
		if isSpace {
			result.WriteString(temp.String())
		} else {
			result.WriteString(temp.String() + " ")
		}
	}
	return strings.TrimSpace(result.String())
}

// Validate checks if the Input contains valid arguments and flags.
func (i *Input) Validate() error {
	if i.Output != "" && !strings.HasSuffix(i.Output, ".txt") {
		return fmt.Errorf("output")
	}
	return nil
}

// ParseArgs parses command-line arguments and sets Input fields.
func (i *Input) ParseArgs() {
	for j, input := range i.Args {
		if j == 0 && i.IsValidFlag(strings.Split(input, "=")[0]) {
			i.Flags[strings.Split(input, "=")[0]] = strings.Split(input, "=")[1]
		}
	}

	if err := i.Validate(); err != nil {
		i.ErrorHandler(err.Error())
	}
}

// IsValidFlag checks if a flag is valid
func (i *Input) IsValidFlag(flag string) bool {
	return validFlags[flag]
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

// RemoveLeadingDashes removes leading '--' from the given string
func (i *Input) RemoveLeadingDashes(input string) string {
	re, _ := regexp.Compile("^--")
	result := re.ReplaceAllString(input, "")

	return result
}
