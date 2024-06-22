package utils

import "flag"

type Input struct {
	Color   string
	Banner  string
	Justify string
	Output  string
	Input   string
	Args    []string
}

var Inputs Input

func init() {
	flag.StringVar(&Inputs.Color, "color", "", "specify a color")
	flag.StringVar(&Inputs.Banner, "banner", "", "specify a banner style")
	flag.StringVar(&Inputs.Justify, "justify", "", "specify text justification")
	flag.StringVar(&Inputs.Output, "output", "", "specify output file")
	flag.StringVar(&Inputs.Input, "input", "", "specify your text")

	// Parse the flags
	flag.Parse()
}
