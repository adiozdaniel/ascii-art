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
	if utils.Inputs.Justify != "" {
		fmt.Printf("🙄 alignment request: 'align=%s'; was ignored\n=================================================\n\n", utils.Inputs.Justify)
	}
	utils.LogOutput(output)
	fmt.Printf("😋 writing.... '%s'. To check output, kindly use: `cat %s | cat -e`\n====================================\n %s", utils.Inputs.Input, utils.Inputs.Output, nonAsciis)
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
func justified() {
	inputChan := make(chan string)
	var ascii_map = ascii.AsciiMap(fileContents)
	prevWidth := 0
	temp := ""

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
				fmt.Println("\n🤩 You were wonderful. Hope you enjoyed.\nExiting the Ascii-Art...")
				os.Exit(0)
			} else {
				temp += input
				scanInput(input)
			}
		default:
			width := utils.GetTerminalWidth()
			if width != prevWidth || temp != "" {
				fmt.Print("\033[H", "\033[2J", "\033[3J", "\033[?25h")
				utils.Alignment(fileContents, ascii_map, output, nonAsciis, width)
				prevWidth = width
				temp = ""
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
