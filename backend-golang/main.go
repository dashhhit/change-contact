package main

import (
	"ChangeContactWeb/cors"
	"ChangeContactWeb/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	corsSettings := cors.Setting()
	app.Use(corsSettings)
	//generateToken.GenerateSecureToken()
	app.Listen(":8000")

}
