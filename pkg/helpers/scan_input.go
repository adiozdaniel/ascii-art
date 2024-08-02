package helpers

// scanInput reads input from CLI interface and updates the input struct.
func ScanInput(input string) {
	cleanInput := utils.RemoveQuotes(input)
	words := strings.Fields(cleanInput)
	var newInput string

	for i, word := range words {
		switch {
		case strings.Contains(word, "--align") || strings.HasPrefix(word, "-align"):
			alignment := strings.TrimPrefix(strings.TrimPrefix(word, "--align="), "-align=")
			if isValidAlignment(alignment) {
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
		case isBannerFile(word):
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
