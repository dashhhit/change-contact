package router

import (
	"ChangeContactWeb/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/auth", handler.Auth)
	app.Post("/getFile", handler.GetFile)
	app.Post("/compareJWT", handler.CompareJWT)
}
