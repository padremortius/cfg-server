package v1

import (
	"github.com/gofiber/fiber/v2"
)

func InitAppRouter(app *fiber.App) {
	app.Get("/:app.:ext", getData)
	app.Get("/:env/:app.:ext", getDataWithEnv)
}
