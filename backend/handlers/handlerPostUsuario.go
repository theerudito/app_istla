package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/theerudito/istla/model/entities"
	"github.com/theerudito/istla/service"
)

type HandlerPostUser struct {
	Service service.IPostUsuario
}

func NewHandlerPostUser(service service.IPostUsuario) *HandlerPostUser {
	return &HandlerPostUser{Service: service}
}

func (cur *HandlerPostUser) GetRegisterByUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"mensaje": "el id es invalido"})
	}

	obj := cur.Service.Get_PostUser_By_UserId(uint(id))

	return c.Status(obj.Codigo).JSON(obj)
}

func (cur *HandlerPostUser) PostRegister(c *fiber.Ctx) error {

	var register entities.PostUsuario

	if err := c.BodyParser(&register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos: " + err.Error(),
		})
	}

	obj := cur.Service.Create_PostUser(register)

	return c.Status(obj.Codigo).JSON(obj)
}

func (cur *HandlerPostUser) PutRegister(c *fiber.Ctx) error {

	var register entities.PostUsuario

	if err := c.BodyParser(&register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos: " + err.Error(),
		})
	}

	obj := cur.Service.Update_PostUser(register)

	return c.Status(obj.Codigo).JSON(obj)
}

func (cur *HandlerPostUser) DeleteRegister(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"mensaje": "el id es invalido"})
	}

	obj := cur.Service.Delete_PostUser(uint(id))

	return c.Status(obj.Codigo).JSON(obj)
}
