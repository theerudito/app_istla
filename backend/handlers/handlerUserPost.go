package handlers

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/theerudito/istla/model/entities"
	"github.com/theerudito/istla/service"
)

type HandlerUserPost struct {
	Service service.IPostUsuario
}

func NewHandlerPostUser(service service.IPostUsuario) *HandlerUserPost {
	return &HandlerUserPost{Service: service}
}

func (cur *HandlerUserPost) GetRegisterByUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"mensaje": "el id es invalido"})
	}

	obj := cur.Service.Get_PostUser_By_UserId(uint(id))

	return c.Status(obj.Codigo).JSON(obj)
}

func (cur *HandlerUserPost) PostRegister(c *fiber.Ctx) error {

	var register entities.PostUsuario

	usuarioIdString, _ := strconv.Atoi(c.FormValue("usuario_id"))
	register.UsuarioId = usuarioIdString
	register.Descripcion = c.FormValue("descripcion")
	register.UsuarioCreacion = c.FormValue("usuario_creacion")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Debe enviar el archivo PDF",
		})
	}

	const maxFileSize = 2 * 1024 * 1024 // 2MB
	if fileHeader.Size > maxFileSize {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El archivo no puede superar los 2MB",
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo abrir el archivo",
		})
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo leer el archivo",
		})
	}

	mimeType := http.DetectContentType(fileBytes[:512])
	if mimeType != "application/pdf" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Solo se permiten archivos PDF",
		})
	}

	if string(fileBytes[:4]) != "%PDF" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "El archivo no es un PDF válido",
		})
	}

	register.File = fileBytes

	obj := cur.Service.Create_PostUser(register)

	return c.Status(obj.Codigo).JSON(obj)
}

func (cur *HandlerUserPost) PutRegister(c *fiber.Ctx) error {

	var register entities.PostUsuario

	usuarioIdString, _ := strconv.Atoi(c.FormValue("usuario_id"))
	postUserIdString, _ := strconv.Atoi(c.FormValue("post_user_id"))
	register.UsuarioId = usuarioIdString
	register.PostUserId = postUserIdString
	register.Descripcion = c.FormValue("descripcion")
	register.UsuarioModificacion = c.FormValue("usuario_modificacion")

	fileHeader, err := c.FormFile("file")
	if err == nil && fileHeader != nil {
		const maxFileSize = 2 * 1024 * 1024 // 2MB
		if fileHeader.Size > maxFileSize {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "El archivo no puede superar los 2MB",
			})
		}

		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "No se pudo abrir el archivo",
			})
		}
		defer file.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "No se pudo leer el archivo",
			})
		}

		mimeType := http.DetectContentType(fileBytes[:512])
		if mimeType != "application/pdf" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Solo se permiten archivos PDF",
			})
		}

		if len(fileBytes) < 4 || string(fileBytes[:4]) != "%PDF" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "El archivo no es un PDF válido",
			})
		}

		register.File = fileBytes
	} else {

		register.File = nil
	}

	obj := cur.Service.Update_PostUser(register)

	return c.Status(obj.Codigo).JSON(obj)
}

func (cur *HandlerUserPost) DeleteRegister(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"mensaje": "el id es invalido"})
	}

	obj := cur.Service.Delete_PostUser(uint(id))

	return c.Status(obj.Codigo).JSON(obj)
}
