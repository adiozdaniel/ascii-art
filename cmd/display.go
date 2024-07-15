package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/adiozdaniel/ascii-art/ascii"
	"github.com/adiozdaniel/ascii-art/routes"
	"github.com/adiozdaniel/ascii-art/utils"
)

// runOutput writes the output to the specified file.
func runOutput() {
	utils.LogOutput(output)
	outputString := fmt.Sprintf("  😋 writing.... '%s'. To check output, kindly use: `cat %s | cat -e` %s", utils.Inputs.Input, utils.Inputs.Output, nonAsciis)
	fmt.Printf("%s\n %s\n", outputString, strings.Repeat("=", len(outputString)-3))
}

// runWeb starts the web server to handle HTTP requests.
func runWeb() {
	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server is running on http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		utils.ErrorHandler("web")
	}
}

// justified runs alignment mode of the application.
// justified runs alignment mode of the application.
func justified() {
	inputChan := make(chan string)
	var inputStr string
	var asciiMap = ascii.AsciiMap(fileContents)
	prevWidth := 0
	var tempInput string
	var prevOutput string

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputChan <- scanner.Text()
		}
		close(inputChan)
	}()

	for {
		select {
		case inputStr = <-inputChan:
			if inputStr == "exit" {
				fmt.Println("\n🤩 You were wonderful. Hope you enjoyed.\nExiting the Ascii-Art...")
				os.Exit(0)
			} else {
				tempInput = inputStr
				scanInput(inputStr)
			}
		default:
			width := utils.GetTerminalWidth()
			var output = ascii.Output(fileContents, tempInput)

			if width != prevWidth || tempInput != "" {
				// Clear only the artwork, not the entire screen
				if prevOutput != "" {
					lines := strings.Split(prevOutput, "\n")
					for range lines {
						fmt.Print("\033[H", "\033[2J", "\033[3J", "\033[?25h")
					}
				}

				termOutput := utils.Alignment(fileContents, asciiMap, output, width)

				// Print new output
				fmt.Print(termOutput)

				// Update previous values
				prevOutput = termOutput
				prevWidth = width
				tempInput = ""
			}

			time.Sleep(2 * time.Second)
		}
	}
}

// scanInput reads input from cli interface and updates the input struct.
func scanInput(input string) {
	utils.Inputs.Args = strings.Split(input, " ")

	// TODO
	// consider adding more options for cli input handling
}
