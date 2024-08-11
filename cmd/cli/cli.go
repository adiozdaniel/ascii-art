package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/adiozdaniel/ascii-art/internals/ascii"
	"github.com/adiozdaniel/ascii-art/internals/models"
	"github.com/adiozdaniel/ascii-art/pkg/helpers"
)

// Global Cli instance
var cli *Cli

// Default values
const (
	defaultFont  = "--standard"
	defaultInput = "Ascii~"
)

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
			"prevWidth": 0,
			"prevColor": "",
			"prevReff":  "",
			"prevFont":  "",
			"tempStr":   "",
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
		flags["font"] != cli.state["prevFont"]
}

// updateDisplay updates the terminal display based on the current state.
func (cli *Cli) updateDisplay(newWidth int) {
	fmt.Println(newWidth)
	flags := cli.app.GetInput().Flags
	banner := cli.app.GetInput().BannerFile[flags["font"]]
	if err := helpers.FileContents(banner); err != nil {
		fmt.Println("Error loading banner file:", err)
	}

	outputs := ascii.Output(flags["input"])
	termOutput := helpers.Alignment(outputs, newWidth)
	// helpers.ClearTerminal()
	fmt.Print(termOutput)
	helpers.ResetCursor()

	cli.state["prevWidth"] = newWidth
	cli.state["tempStr"] = ""
	cli.state["prevColor"] = flags["color"]
	cli.state["prevReff"] = flags["reff"]
	cli.state["prevFont"] = flags["font"]
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

	cli.app.GetInput().Flags["font"] = defaultFont
	cli.app.GetInput().Flags["input"] = defaultInput
}

// main runs the CLI application.
func main() {
	cli.app.GetInput().Init()

	for {
		select {
		case input, ok := <-cli.inputChan:
			if !ok {
				fmt.Println("Input channel closed. Exiting...")
				return
			}
			if input == "exit" {
				fmt.Println("\n🤩 You were wonderful. Hope you enjoyed.\nExiting the Ascii-Art...")
				return
			} else if input != "" {
				cli.state["tempStr"] = input
				helpers.ScanInput(input)
			}
		default:
			newWidth := helpers.GetTerminalWidth()
			if cli.shouldUpdate(newWidth) {
				cli.updateDisplay(newWidth)
			}
			time.Sleep(2 * time.Second)
		}
	}
}
