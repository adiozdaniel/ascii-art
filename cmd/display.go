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

// justified runs the alignment mode of the application.
func justified() {
	inputChan := make(chan string)
	prevWidth := 0
	tempStr := ""

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputChan <- scanner.Text()
		}
		close(inputChan)
	}()

	for {
		select {
		case input := <-inputChan:
			if input == "exit" {
				fmt.Println("\n\n🤩 You were wonderful. Hope you enjoyed.\nExiting the Ascii-Art...")
				return
			} else {
				scanInput(input)
				tempStr = input
			}
		default:
			newWidth := utils.GetTerminalWidth()
			outputs := ascii.Output(tempStr)

			if newWidth != prevWidth || tempStr != "" {
				// fmt.Print("\033[H", "\033[2J", "\033[3J", "\033[?25h")
				termOutput := utils.Alignment(outputs, newWidth)
				fmt.Print(termOutput)
				prevWidth = newWidth
				tempStr = ""
			}
			time.Sleep(2 * time.Second)
		}
	}
}

// scanInput reads input from cli interface and updates the input struct.
func scanInput(input string) {
	utils.Inputs.Input = input

	// TODO
	// consider adding more options for cli input handling
}
