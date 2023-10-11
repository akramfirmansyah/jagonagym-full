package routes

import (
	"github.com/akramfirmansyah/jagona-gym/server/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterMemberRoutes(router fiber.Router) {
	router.Post("/", controllers.CreateMember)
}
