package main

import (
	"os"

	"github.com/akramfirmansyah/jagona-gym/server/config"
	"github.com/akramfirmansyah/jagona-gym/server/controllers"
	"github.com/akramfirmansyah/jagona-gym/server/database"
	"github.com/akramfirmansyah/jagona-gym/server/middlewares"
	"github.com/akramfirmansyah/jagona-gym/server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	config.LoadEnv()
	database.ConnectDB()
	database.Migrate()
	database.Seeder()
}

func main() {
	app := fiber.New(fiber.Config{
		AppName:     "Jagona Gym API v0.0.1",
		JSONEncoder: jsoniter.Marshal,
		JSONDecoder: jsoniter.Unmarshal,
	})

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Post("/login", controllers.Login)

	api := app.Group("/api", middlewares.Protected())

	// Member
	routes.RegisterMemberRoutes(api.Group("/member"))

	app.Listen(":" + os.Getenv("PORT"))
}
