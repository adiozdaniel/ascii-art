package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	reverse "github.com/adiozdaniel/ascii-art/cmd/rev"
	"github.com/adiozdaniel/ascii-art/internals/ascii"
	"github.com/adiozdaniel/ascii-art/internals/models"
	"github.com/adiozdaniel/ascii-art/pkg/helpers"
)

// Global Cli instance
var cli *Cli

// Cli holds the state for the CLI interface.
type Cli struct {
	app       *models.StateManager
	inputChan chan string
	state     map[string]interface{}
}

// NewCli creates a new Cli instance.
func NewCli(sm *models.StateManager) *Cli {
	return &Cli{
		app:       sm,
		inputChan: make(chan string),
		state: map[string]interface{}{
			"prevWidth":   0,
			"prevColor":   "",
			"prevReff":    "",
			"prevFont":    "",
			"tempStr":     "",
			"prevReverse": "",
		},
	}
}

// readInput reads input from the CLI interface and sends it to the input channel.
func (cli *Cli) readInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cli.inputChan <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
	close(cli.inputChan)
}

// shouldUpdate checks if the terminal output needs to be updated.
func (cli *Cli) shouldUpdate(newWidth int) bool {
	flags := cli.app.GetInput().Flags
	return newWidth != cli.state["prevWidth"].(int) || cli.state["tempStr"] != "" ||
		flags["color"] != cli.state["prevColor"] ||
		flags["reff"] != cli.state["prevReff"] ||
		flags["font"] != cli.state["prevFont"] || flags["reverse"] != cli.state["prevReverse"]
}

// updateDisplay updates the terminal display based on the current state.
func (cli *Cli) updateDisplay(newWidth int) {
	flags := cli.app.GetInput().Flags
	banner := cli.app.GetInput().BannerFile[flags["font"]]
	if err := helpers.FileContents(banner); err != nil {
		fmt.Println("Error loading banner file:", err)
	}

	outputs := ascii.Output(flags["input"])
	termOutput := helpers.Alignment(outputs, newWidth)
	helpers.ClearTerminal()
	if cli.app.GetInput().Flags["output"] != "" {
		ascii.LogOutput(termOutput)
		cli.state["prevWidth"] = newWidth
		cli.state["tempStr"] = ""
		cli.state["prevReverse"] = ""
		cli.state["prevColor"] = flags["color"]
		cli.state["prevReff"] = flags["reff"]
		cli.state["prevFont"] = flags["font"]
		cli.app.GetInput().Flags["output"] = ""
		cli.app.GetInput().Flags["reverse"] = ""
		return
	}

	if cli.app.GetInput().Flags["reverse"] != "" {
		reverse.CheckReverse(flags["reverse"])
		cli.state["prevWidth"] = newWidth
		cli.state["tempStr"] = ""
		cli.state["prevReverse"] = ""
		cli.state["prevColor"] = flags["color"]
		cli.state["prevReff"] = flags["reff"]
		cli.state["prevFont"] = flags["font"]
		cli.app.GetInput().Flags["output"] = ""
		cli.app.GetInput().Flags["reverse"] = ""
		return
	}

	fmt.Print("\n", termOutput)

	Footer()

	cli.state["prevWidth"] = newWidth
	cli.state["tempStr"] = ""
	cli.state["prevColor"] = flags["color"]
	cli.state["prevReff"] = flags["reff"]
	cli.state["prevFont"] = flags["font"]
}

// Footer writes the footer for CLI mode
func Footer() {
	var title = "ASCII Art Application - CLI Mode. Version 5.0.1"
	var message = "Type 'exit' to quit"
	var help = "For help, use the documentation"
	var _, col = helpers.GetTerminalWidth()
	var footer = fmt.Sprintf("\n%s\n%s\n%s\n%s\n%s\n%s\n",
		strings.Repeat("=", col), ascii.NonAsciiOutput(),
		title, message, help,
		strings.Repeat("=", col),
	)
	fmt.Printf("%s", strings.Repeat("\n", 5))
	fmt.Println(footer)
	time.Sleep(1 * time.Second)
	fmt.Printf("\033[%dA", 9)
}

// init initializes the CLI interface.
func init() {
	cli = NewCli(models.GetStateManager())

	bc, err := cli.app.GetConfig().CreateBannerCache()
	if err != nil {
		cli.app.GetInput().ErrorHandler("banners")
	}

	go cli.readInput()
	cli.app.GetConfig().BannerFileCache = bc
}

// runCli runs the CLI application.
func runCli() error {
	for {
		select {
		case input, ok := <-cli.inputChan:
			if !ok {
				fmt.Println("Input channel closed. Exiting...")
				return nil
			}
			if input == "exit" {
				helpers.ClearTerminal()
				fmt.Println("\n🤩 You were wonderful. Hope you enjoyed.\nExiting the Ascii-Art...")
				return nil
			} else if input != "" {
				cli.state["tempStr"] = input
				helpers.ScanInput(input)
			}
		default:
			_, newWidth := helpers.GetTerminalWidth()
			if cli.shouldUpdate(newWidth) {
				cli.updateDisplay(newWidth)
			}
			time.Sleep(2 * time.Second)
		}
	}
}

// main runs the CLI application.
func main() {
	cli.app.GetInput().Init()

	if err := runCli(); err != nil {
		cli.app.GetInput().ErrorHandler("cli")
	}
}
