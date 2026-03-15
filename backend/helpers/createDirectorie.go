package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFolder() error {

	directorioPrincipal := "resources"

	subdirectorios := []string{
		"pdf", "imagen",
	}

	if err := os.MkdirAll(directorioPrincipal, os.ModePerm); err != nil {
		return fmt.Errorf("error al crear el directorio principal: %w", err)
	}

	for _, sub := range subdirectorios {
		path := filepath.Join(directorioPrincipal, sub)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("error al crear el subdirectorio %s: %w", sub, err)
		}
	}

	return nil
}
