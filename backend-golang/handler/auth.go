package handler

import (
	"ChangeContactWeb/db"
	"ChangeContactWeb/model"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	token := new(model.Token)
	c.Set("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
	if err := c.BodyParser(token); err != nil {

		return c.JSON(model.Response{
			Data:    "",
			Message: err.Error(),
			Code:    403,
		})
	}
	if err := db.Compare(token); err != nil {
		return c.JSON(model.Response{
			Data:    "",
			Message: err.Error(),
			Code:    403,
		})
	}
	return c.JSON(model.Response{
		Data:    token.Token,
		Message: "",
		Code:    200,
	})
}
