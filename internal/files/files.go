package files

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

func StatFile(path string) error {
	info, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("target file does not exist")
	}

	if err != nil {
		return fmt.Errorf("failed to check %s file: %v", path, err)
	}

	if info.IsDir() {
		return fmt.Errorf("target file path is a directory")
	}

	return nil
}

func ReadYAML(path string, data any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %v", path, err)
	}

	if err := yaml.Unmarshal(b, data); err != nil {
		return fmt.Errorf("failed to marshal dump schema: %v", err)
	}

	return nil
}

func WriteYAML(path string, data any) error {
	var b bytes.Buffer
	enc := yaml.NewEncoder(&b)
	enc.SetIndent(2)

	if err := enc.Encode(&data); err != nil {
		return fmt.Errorf("failed to marshal dump schema: %v", err)
	}

	if err := os.WriteFile(path, b.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed open file %s for writing: %v", path, err)
	}

	return nil
}
