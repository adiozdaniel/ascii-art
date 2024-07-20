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
	prevColor := ""
	prevReff := ""
	prevBanner := ""
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
				fmt.Println("\n🤩 You were wonderful. Hope you enjoyed.\nExiting the Ascii-Art...")
				os.Exit(0)
			} else {
				if input != "" {
					tempStr = input
					scanInput(input)
				}
			}
		default:
			newWidth := utils.GetTerminalWidth()
			if newWidth != prevWidth || tempStr != "" || utils.Inputs.Color != prevColor || utils.Inputs.ColorRef != prevReff || utils.Inputs.BannerPath != prevBanner {
				outputs := ascii.Output(utils.Inputs.Input)
				termOutput := utils.Alignment(outputs, newWidth)
				fmt.Print("\033[H", "\033[2J", "\033[3J", "\033[?25h")
				fmt.Print(termOutput)
				fmt.Print("\033[999;1H")
				prevWidth = newWidth
				tempStr = ""
				prevColor = utils.Inputs.Color
				prevReff = utils.Inputs.ColorRef
				prevBanner = utils.Inputs.BannerPath
			}
			time.Sleep(2 * time.Second)
		}
	}
}

// scanInput reads input from CLI interface and updates the input struct.
func scanInput(input string) {
	cleanInput := utils.RemoveQuotes(input)
	words := strings.Fields(cleanInput)
	newInput := ""

	for i, word := range words {
		switch {
		case strings.Contains(word, "--align") || strings.HasPrefix(word, "-align"):
			alignment := strings.TrimPrefix(strings.TrimPrefix(word, "--align="), "-align=")
			if alignment == "left" || alignment == "center" || alignment == "right" || alignment == "justify" {
				utils.Inputs.Justify = alignment
				continue
			}
			utils.ErrorHandler("justify")
		case strings.Contains(word, "--color") || strings.Contains(word, "-color"):
			color := strings.TrimPrefix(strings.TrimPrefix(word, "--color="), "-color=")
			if color != "" {
				utils.Inputs.Color = color
				continue
			}
			utils.ErrorHandler("colors")
		case strings.Contains(word, "--reff") || strings.Contains(word, "-reff"):
			reff := strings.TrimPrefix(strings.TrimPrefix(word, "--reff="), "-reff=")
			if reff != "" {
				utils.Inputs.ColorRef = reff
				continue
			}
			utils.ErrorHandler("colors")
		case strings.Contains(word, "--output") || strings.Contains(word, "-output"):
			fmt.Println("🙄 Sorry, FS Mode cannot be set in alignment mode.")
			os.Exit(0)
		case strings.Contains(word, "--standard") || strings.Contains(word, "--thinkertoy") || strings.Contains(word, "--shadow"):
			if value, ok := utils.BannerFiles[word]; ok {
				utils.Inputs.BannerPath = value
				continue
			}

			newInput += word + " "
		case strings.Contains(word, "standard") || strings.Contains(word, "thinkertoy") || strings.Contains(word, "shadow"):
			if i == len(words)-1 && len(words) != 1 {
				if value, ok := utils.BannerFiles[word]; ok {
					utils.Inputs.BannerPath = value
				}
				break
			}
			newInput += word + " "
		default:
			newInput += word + " "
		}
	}

	if newInput != "" {
		utils.Inputs.Input = strings.TrimSpace(newInput)
	}
}
