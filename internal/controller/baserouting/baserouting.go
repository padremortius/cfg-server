package baserouting

import (
	"cfg-server/internal/config"
	"cfg-server/internal/controller/structs"

	fiber "github.com/gofiber/fiber/v2"
)

func getHealth(c *fiber.Ctx) error {
	return c.JSON(structs.Health{Status: "up"})
}

func getInfo(c *fiber.Ctx) error {
	return c.JSON(config.Cfg.Version)
}

func getEnv(c *fiber.Ctx) error {
	return c.JSON(config.Cfg)
}

func InitBaseRouter(app *fiber.App) {
	// K8s probe
	app.Get("/health", getHealth)

	//info about service
	app.Get("/info", getInfo)

	//env
	app.Get("/env", getEnv)
}
