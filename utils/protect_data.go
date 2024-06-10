package utils

import (
	"os"
	"path/filepath"
)

func ProtectFilesInDirectory() error {
	directoryPath := "../banners"

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Protect the file
			err := ProtectFile(path)
			if err != nil {
				return err
			} else {
				return nil
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
