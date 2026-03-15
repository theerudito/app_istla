package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theerudito/istla/model/dto"
	"github.com/theerudito/istla/model/entities"
	"github.com/theerudito/istla/service"
)

type HandlerUser struct {
	Service service.IUser
}

func NewHandlerUser(service service.IUser) *HandlerUser {
	return &HandlerUser{Service: service}
}

func (cu *HandlerUser) Login(c *fiber.Ctx) error {

	var login dto.UsuarioLoginDTO

	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos: " + err.Error(),
		})
	}

	obj := cu.Service.Login(login)

	return c.Status(obj.Codigo).JSON(obj)
}

func (cu *HandlerUser) Register(c *fiber.Ctx) error {

	var register entities.Usuario

	if err := c.BodyParser(&register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos: " + err.Error(),
		})
	}

	obj := cu.Service.Register(register)

	return c.Status(obj.Codigo).JSON(obj)
}
