package handler

import (
	"ChangeContactWeb/db"
	"ChangeContactWeb/model"
	"github.com/gofiber/fiber/v2"
)

func GetFile(c *fiber.Ctx) error {
	authorizationToken := c.Get("Authorization")[7:]
	c.Set("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
	if err := db.Compare(&model.Token{Token: authorizationToken}); err != nil {
		return c.Status(403).JSON(model.Response{
			Data:    nil,
			Message: err.Error(),
			Code:    400,
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(403).JSON(model.Response{
			Data:    nil,
			Message: err.Error(),
			Code:    400,
		})
	}

	file2, err := file.Open()
	defer file2.Close()
	if err != nil {
		return c.Status(403).JSON(model.Response{
			Data:    nil,
			Message: err.Error(),
			Code:    400,
		})
	}

	//fileOs, ok := file2.(*os.File)
	//if !ok {
	//	return c.JSON(model.Response{
	//		Data:    nil,
	//		Message: "Invalid file",
	//		Code:    400,
	//	})
	//}
	correctFile, err := ChangeContactsMultiPart(file2)
	if err != nil {
		return c.Status(403).JSON(model.Response{
			Data:    nil,
			Message: err.Error(),
			Code:    400,
		})
	}
	err = db.UpdateCount(&model.Token{Token: authorizationToken})
	if err != nil {
		return c.Status(403).JSON(model.Response{
			Data:    nil,
			Message: err.Error(),
			Code:    400,
		})
	}

	return c.SendStream(correctFile)
}
