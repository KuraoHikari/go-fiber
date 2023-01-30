package routes

import (
	"go-fiber/config"
	"go-fiber/handler"
	"go-fiber/middleware"
	"go-fiber/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(route *fiber.App) {
	route.Static("/public", config.ProjectRootPath+"/public/assets")

	route.Post("/login", handler.Login)

	route.Get("/user", middleware.Auth, handler.UserHandlerGetAll)
	route.Get("/user/:id", handler.UserHandlerGetById)
	route.Post("/user", handler.UserHandlerCreate)
	route.Put("user/:id", handler.UserHandlerUpdate)
	route.Delete("user/:id", handler.UserHandlerDelete)

	route.Post("/book", utils.HandleSingleFile, handler.BookHandlerCreate)
}
