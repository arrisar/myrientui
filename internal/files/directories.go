package files

import (
	"errors"
	"fmt"
	"os"
)

func StatDir(path string) error {
	info, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("target directory does not exist")
	}

	if err != nil {
		return fmt.Errorf("failed to check %s directory: %v", path, err)
	}

	if !info.IsDir() {
		return fmt.Errorf("target directory path is a file")
	}

	return nil
}

func MkDir(path string) error {
	if err := os.MkdirAll(path, 0777); err != nil {
		return fmt.Errorf("failed to setup %s directory: %v", path, err)
	}

	return nil
}
