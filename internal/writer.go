package internal

import (
	"fmt"
	"os"
)

// Write to output file
func Write(path, txt string) error {
	if txt == "" { // skip to create file because of empty
		return nil
	}

	// Create file
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file error: %w", err)
	}
	defer file.Close()

	// Write to file
	_, err = file.WriteString(txt)
	if err != nil {
		return fmt.Errorf("write file error: %w", err)
	}

	return nil
}
