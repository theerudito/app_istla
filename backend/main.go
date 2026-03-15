package main

import (
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

	// CREAR INSTANCIA DE FIBER
	app := fiber.New()

	// LEER LAS VARIABLES DE ENTORNO
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

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

	// INYECTAR CONTROLADORES
	handlerUser := handlers.NewHandlerUser(repositorieUser)
	handlerUseregister := handlers.NewHandlerPostUser(repositorieUserRegister)

	// REGISTRAR LOS CONTROLADORES
	controllers := &handlers.Handlers{
		User:         handlerUser,
		UserRegister: handlerUseregister,
	}

	router.SetupRoutes(app, controllers)

	// LEVANTAR EL SERVIDOR
	log.Println("CREADO POR BETWEEN BYTES SOFTWARE")
	log.Fatal(app.Listen(":" + os.Getenv("PortServer")))

}
