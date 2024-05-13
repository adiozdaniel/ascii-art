package utilities

import (
	"os"
)

func ProtectFile(filename string) error {
	// Open the file with read-only permission
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get the current file permissions
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Set the read-only permission for the owner, group, and others
	err = os.Chmod(filename, info.Mode().Perm()&0444)
	if err != nil {
		return err
	}

	return nil
}
