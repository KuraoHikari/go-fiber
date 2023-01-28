package routes

import (
	"go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(route *fiber.App) {
	route.Get("/user", handler.UserHandlerGetAll)
	route.Get("/user/:id", handler.UserHandlerGetById)
	route.Post("/user", handler.UserHandlerCreate)
	route.Put("user/:id", handler.UserHandlerUpdate)
	route.Delete("user/:id", handler.UserHandlerDelete)
}
