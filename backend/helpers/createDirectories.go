package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFolder() error {

	sourcePath := os.Getenv("Source_Path")
	pdfFolder := os.Getenv("PDF")
	imgFolder := os.Getenv("IMAGEN")

	if sourcePath == "" || pdfFolder == "" || imgFolder == "" {
		return fmt.Errorf("error: Source_Path, PDF o IMAGEN no están definidos en las variables de entorno")
	}

	subdirectorios := []string{pdfFolder, imgFolder}

	if err := os.MkdirAll(sourcePath, os.ModePerm); err != nil {
		return fmt.Errorf("error al crear el directorio principal (%s): %w", sourcePath, err)
	}

	for _, sub := range subdirectorios {
		path := filepath.Join(sourcePath, sub)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("error al crear el subdirectorio %s: %w", sub, err)
		}
	}

	return nil
}
