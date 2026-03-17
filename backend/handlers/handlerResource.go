package handlers

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func ResourceController(c *fiber.Ctx) error {

	root := os.Getenv("Source_Path")
	if root == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Source_Path no definido"})
	}

	fileName := c.Params("file")
	folder := c.Params("folder")

	var filePath string

	switch folder {
	case "pdf":
		filePath = filepath.Join(root, os.Getenv("PDF"), fileName)
	case "imagen":
		filePath = filepath.Join(root, os.Getenv("IMAGEN"), fileName)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "folder inválido"})
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "archivo no encontrado"})
	}

	return c.SendFile(filePath)
}
