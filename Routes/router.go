package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/julienb86/Api_Social_Network/Controllers"
)

func PostRoutes(route fiber.Router) {
	route.Get("", controllers.GetAllPosts)

}
