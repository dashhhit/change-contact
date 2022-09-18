package cors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setting() fiber.Handler {
	c := cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:3000, http://127.0.0.1, http://localhost:3000",
		AllowCredentials: true,
		AllowHeaders:     "Access-Control-Request-Headers, Access-Control-Request-Method, Host, Connection, Referer, Sec-Fetch-Dest, Sec-Fetch-Site, Sec-Fetch-Mode, Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Allow-Origin, token, User-Agent",
		ExposeHeaders:    "Access-Control-Request-Headers, Access-Control-Request-Method, Host, Connection, Referer, Sec-Fetch-Dest, Sec-Fetch-Site, Sec-Fetch-Mode, Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Allow-Origin, token, User-Agent",
	})
	return c
}

// Insert the middleware
