package v1

import (
	"github.com/gofiber/fiber/v2"
)

func InitAppRouter(app *fiber.App) {
	app.Get("/:app.json", getDataInJSON)
	app.Get("/:app.yaml", getDataInYaml)
	app.Get("/:app.yml", getDataInYaml)
	app.Get("/:env/:app.json", getDataWithEnvInJSON)
	app.Get("/:env/:app.yaml", getDataWithEnvInYaml)
	app.Get("/:env/:app.yml", getDataWithEnvInYaml)
}
