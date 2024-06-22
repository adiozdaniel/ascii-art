package utils

// GetFile returns the ascii graphic filepath to use.
func getFile() {
	ourBanner := "../banners/standard.txt"
	flag := ""

	for _, val := range Inputs.Args {
		if value, ok := bannerFiles[val]; ok {
			ourBanner = value
			flag = val
		}
	}

	Inputs.Banner = ourBanner
	if flag == "" {
		Inputs.isBanner = false
	}
}
