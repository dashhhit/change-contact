package handler

import (
	"ChangeContactWeb/db"
	"ChangeContactWeb/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func CompareJWT(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
	authorizationToken := c.Get("Authorization")[7:]
	fmt.Println(authorizationToken)
	if err := db.Compare(&model.Token{Token: authorizationToken}); err != nil {
		return c.JSON(model.Response{
			Data:    "",
			Code:    403,
			Message: "Wrong token",
		})
	}
	return c.JSON(model.Response{
		Data:    "",
		Code:    200,
		Message: "",
	})
}
