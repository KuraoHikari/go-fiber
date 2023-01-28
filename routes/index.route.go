package routes

import (
	"go-fiber/config"
	"go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func middleware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token != "secret" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	return ctx.Next()
}

func RouteInit(route *fiber.App) {
	route.Static("/public", config.ProjectRootPath+"/public/assets")

	route.Get("/user", middleware, handler.UserHandlerGetAll)
	route.Get("/user/:id", handler.UserHandlerGetById)
	route.Post("/user", handler.UserHandlerCreate)
	route.Put("user/:id", handler.UserHandlerUpdate)
	route.Delete("user/:id", handler.UserHandlerDelete)
}
