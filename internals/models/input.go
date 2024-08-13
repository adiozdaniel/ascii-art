package models

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// InputData holds all the data passed around in the application
type InputData struct {
	Flags        map[string]string
	BannerFile   map[string]string
	ValidFlags   map[string]bool
	Args         []string
	FileContents []string
}

// NewInputData creates a new InputData instance with default values
func NewInputData() *InputData {
	return &InputData{
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

// Init initializes the InputData
func (i *InputData) Init() {
	if len(os.Args) > 1 && strings.Contains(os.Args[1], "-test") {
		return
	}

	if len(os.Args) > 1 {
		i.ErrorHandler("fatal")
	}
}

// BannerFiles returns the map of banner files for the InputData instance
func (i *InputData) BannerFiles() map[string]string {
	return i.BannerFile
}

// GetBannerPath returns the name for a specific banner key or defaults to "standard.txt"
func (i *InputData) GetBannerPath(key string) string {
	if name, ok := i.BannerFile[key]; ok {
		return name
	}
	return "standard.txt"
}

// Validate checks if the InputData contains valid arguments and flags
func (i *InputData) Validate() error {
	if i.Flags["output"] != "" && !strings.HasSuffix(i.Flags["output"], ".txt") {
		return fmt.Errorf("output file must have a .txt extension")
	}
	return nil
}

// ParseArgs parses command-line arguments and sets InputData fields
func (i *InputData) ParseArgs() {
	if i.Flags == nil {
		i.Flags = make(map[string]string)
	}

	i.Checkbanner()
	j := 0
	for j < len(i.Args){
		i.CheckReff(i.Args[j], j)
		j++
	}

	if len(i.Args) > 0 {
		i.Flags["input"] = strings.Join(i.Args, " ")
	}
	i.Args = nil

	if i.Flags["color"] != "" && i.Flags["reff"] == "" {
		i.Flags["reff"] = i.Flags["input"]
	}
}

// IsValidFlag checks if a flag is valid
func (i *InputData) IsValidFlag(flag string) bool {
	return i.ValidFlags[flag]
}

// Checkbanner checks if arguments has bannerfile
func (i *InputData) Checkbanner() {
	if len(i.Args) == 1 || len(i.Args) == 0 {
		return
	}
	if _, ok := bannerFiles[i.Args[len(i.Args)-1]]; ok {
		i.Flags["font"] = i.Args[len(i.Args)-1]
		i.Args = i.Args[:len(i.Args)-1]
	}
}

func (i *InputData) CheckReff(flag string, j int) {
	parts := strings.Split(flag, "=")

	if len(parts) == 2 {
		flag := parts[0]
		value := parts[1]
		fmt.Println(flag, value)
		parsedFlag := i.RemoveLeadingDashes(flag)
		i.Flags[parsedFlag] = value
		i.Args = i.Args[j+1:]
	}
}

// RemoveLeadingDashes removes leading '--' from the given string
func (i *InputData) RemoveLeadingDashes(InputData string) string {
	if strings.HasPrefix(InputData, "--") {
		return InputData[2:]
	}
	return InputData
}

// RemoveQuotes removes opening or closing quotes in a string
func (i *InputData) RemoveQuotes(InputData string) string {
	var result strings.Builder
	inQuotes := false

	for _, ch := range InputData {
		if ch == '"' {
			inQuotes = !inQuotes
		} else if ch != '\'' {
			result.WriteRune(ch)
		}
	}

	return strings.TrimSpace(result.String())
}

// GetProjectRoot dynamically finds the project root directory
func (i *InputData) GetProjectRoot(path, name string) string {
	cwd, _ := os.Getwd()
	baseDir := cwd
	if strings.HasSuffix(baseDir, "/web") || strings.HasSuffix(baseDir, "/cli") {
		baseDir = filepath.Join(cwd, "../../")
	}
	return filepath.Join(baseDir, path, name)
}
