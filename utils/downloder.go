package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(url string, bannerPath string) error {
	dir := filepath.Dir(bannerPath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	// Create the file
	file, err := os.Create(bannerPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("🧐 oops, check your internet connection")
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
