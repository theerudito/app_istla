package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func SaveImageToDirectory(file []byte, name string, ext string, folder string) (string, error) {
	if len(file) == 0 {
		return "", fmt.Errorf("el archivo está vacío")
	}

	root := os.Getenv("Source_Path")
	baseURL := strings.TrimRight(os.Getenv("Url"), "/")

	if root == "" || baseURL == "" {
		return "", fmt.Errorf("Source_Path o Url no están definidos (Root: %s, Base: %s)", root, baseURL)
	}

	if folder == "" {
		return "", fmt.Errorf("el parámetro folder llegó vacío")
	}

	dir := filepath.Join(root, folder)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", fmt.Errorf("error creando carpeta %s: %w", dir, err)
	}

	fileName := name + ext
	filePath := filepath.Join(dir, fileName)

	err := os.WriteFile(filePath, file, 0644)
	if err != nil {
		return "", fmt.Errorf("error al guardar archivo en %s: %w", filePath, err)
	}

	publicURL := fmt.Sprintf("%s/%s/%s/%s", baseURL, root, folder, fileName)

	return publicURL, nil
}

func DeleteImageFromDirectory(folder string, fileName string) error {
	root := os.Getenv("Source_Path")
	if root == "" {
		return fmt.Errorf("error: la variable Source_Path no está definida en el entorno")
	}

	filePath := filepath.Join(root, folder, fileName)
	log.Printf("Intentando eliminar archivo en la ruta física: %s", filePath)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("El archivo no existe físicamente: %s", filePath)
		return nil
	}

	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("no se pudo eliminar el archivo %s: %w", filePath, err)
	}

	log.Println("Archivo eliminado exitosamente:", filePath)
	return nil
}
