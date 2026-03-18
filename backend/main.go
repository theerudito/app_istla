package main

import (
	"errors"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/theerudito/istla/db"
	"github.com/theerudito/istla/handlers"
	"github.com/theerudito/istla/helpers"
	"github.com/theerudito/istla/repositories"
	"github.com/theerudito/istla/router"
)

func main() {

	// LEER LAS VARIABLES DE ENTORNO
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: No se encontró el archivo .env, usando variables de entorno del sistema")
	}

	// CREAR INSTANCIA DE FIBER
	app := fiber.New(fiber.Config{
		BodyLimit: 2 * 1024 * 1024, // 2MB
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			switch code {

			case fiber.StatusRequestEntityTooLarge:
				return c.Status(code).JSON(fiber.Map{
					"error": "El archivo excede el tamaño máximo permitido (2MB)",
				})

			default:
				return c.Status(code).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
		},
	})

	// CREAR LAS CARPETAS
	er := helpers.CreateFolder()
	if er != nil {
		log.Fatal(er)
	}

	// CONECTAR A LA DB
	_db, err := db.ConnectarDB()

	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	} else {
		log.Println("connectada la db en ", "postgres")
	}

	defer _db.Close()

	conn := _db.GetConn()

	// INYECTAR REPOSITORIOS
	repositorieUser := repositories.NewRepositorieUser(conn)
	repositorieUserRegister := repositories.NewRepositoriePostUser(conn)
	repositorieProfile := repositories.NewRepositoriePerfil(conn)

	// INYECTAR CONTROLADORES
	handlerUser := handlers.NewHandlerUser(repositorieUser)
	handlerUseregister := handlers.NewHandlerPostUser(repositorieUserRegister)
	handlerProfile := handlers.NewHandlerPerfil(repositorieProfile)

	// REGISTRAR LOS CONTROLADORES
	controllers := &handlers.HandlersRegister{
		User:         handlerUser,
		UserRegister: handlerUseregister,
		Profile:      handlerProfile,
	}

	router.SetupRoutes(app, controllers)

	// LEVANTAR EL SERVIDOR
	log.Println("CREADO POR BETWEEN BYTES SOFTWARE")
	log.Fatal(app.Listen(":" + os.Getenv("PortServer")))

}
