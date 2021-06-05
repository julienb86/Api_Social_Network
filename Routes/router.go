package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/julienb86/Api_Social_Network/controllers"
)

func PostRoutes(route fiber.Router) {
	route.Get("", controllers.GetAllPosts)

}