package handlers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func ResourceController(c *fiber.Ctx) error {
	root := os.Getenv("Source_Path")
	if root == "" {
		root = "resources"
	}

	fileName := c.Params("file")
	folderParam := c.Params("folder")

	pdfDir := os.Getenv("PDF")
	imgDir := os.Getenv("IMAGEN")

	var filePath string

	switch folderParam {
	case "pdf", pdfDir:

		if pdfDir == "" {
			pdfDir = "pdf"
		}
		filePath = filepath.Join(root, pdfDir, fileName)

	case "imagen", imgDir:
		if imgDir == "" {
			imgDir = "imagen"
		}
		filePath = filepath.Join(root, imgDir, fileName)

	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Folder '%s' no es válido", folderParam),
		})
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "El archivo no existe en la ruta: " + filePath,
		})
	}

	return c.SendFile(filePath)
}
