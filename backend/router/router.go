package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/theerudito/istla/handlers"
)

func SetupRoutes(app *fiber.App, controller *handlers.HandlersRegister) {
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

	v1.Get("/profiles", controller.Profile.GetProfiles)

	v1.Post("/login", controller.User.Login)
	v1.Post("/register", controller.User.Register)

	v1.Get("/post/get_by_user/:id", controller.UserPost.GetRegisterByUser)
	v1.Post("/post", controller.UserPost.PostRegister)
	v1.Put("/post", controller.UserPost.PutRegister)
	v1.Delete("/post/:id", controller.UserPost.DeleteRegister)

	v1.Get("/resources/:folder/:file", handlers.ResourceController)
}
