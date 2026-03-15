package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/theerudito/istla/handlers"
)

func SetupRoutes(app *fiber.App, handlers *handlers.Handlers) {
	allowedOrigins := map[string]bool{
		os.Getenv("URL_Frontend"): true,
	}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS, PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowOriginsFunc: func(origin string) bool {
			return allowedOrigins[origin]
		},
	}))

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/login", handlers.User.Login)
	v1.Post("/register", handlers.User.Register)

	v1.Get("/post/get_by_user/id:", handlers.UserRegister.GetRegisterByUser)
	v1.Post("/post", handlers.UserRegister.PostRegister)
	v1.Put("/post", handlers.UserRegister.PutRegister)
	v1.Delete("/post/:id", handlers.UserRegister.DeleteRegister)
}
