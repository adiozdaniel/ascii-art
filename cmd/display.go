package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/adiozdaniel/ascii-art/routes"
	"github.com/adiozdaniel/ascii-art/utils"
)

// runOutput writes the output to the specified file.
func runOutput(output, nonAsciis string) {
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

func justified(output, nonAsciis string) {
	inputChan := make(chan string)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputChan <- scanner.Text()
		}
	}()

	for {
		select {
		case input := <-inputChan:
			if input == "exit" {
				fmt.Println("Exiting...")
				return
			} else {
				fmt.Printf("Received input: %s\n", input)
			}
		default:
			width := utils.GetTerminalWidth()
			utils.Alignment(output, nonAsciis, width)

			time.Sleep(2 * time.Second)

			// fmt.Print("\033[H\033[2J")

			fmt.Println("To exit, type \"exit\" at the bottom of the screen.")
		}
	}
}
