package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

// DownloadFile downloads the banner file and/or creates the banner directory
func DownloadFile(url string, bannerPath string) error {
	dir := filepath.Dir(bannerPath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(bannerPath)
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("🧐 oops, check your internet connection")
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// CleanPath removes /cmd/ from a path to clean absolute paths
func CleanPath(path, dirSeg string) string {
	escapedDirSegment := regexp.QuoteMeta(dirSeg)
	pattern := fmt.Sprintf(`%s\/`, escapedDirSegment)
	regex := regexp.MustCompile(pattern)
	newPath := regex.ReplaceAllString(path, "")

	return newPath
}
