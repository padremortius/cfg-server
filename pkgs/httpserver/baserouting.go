package httpserver

import (
	fiber "github.com/gofiber/fiber/v2"
)

type (
	BaseRoutes struct {
		cfg     any
		version any
	}

	Health struct {
		Status string
	}
)

func (b *BaseRoutes) getHealth(c *fiber.Ctx) error {
	return c.JSON(Health{Status: "up"})
}

func (b *BaseRoutes) getInfo(c *fiber.Ctx) error {
	return c.JSON(b.version)
}

func (b *BaseRoutes) getEnv(c *fiber.Ctx) error {
	return c.JSON(b.cfg)
}

func InitBaseRouter(app *fiber.App, aConfig any, aVersion any) {
	bRoutes := BaseRoutes{cfg: aConfig, version: aVersion}

	// K8s probe
	app.Get("/health", bRoutes.getHealth)

	//info about service
	app.Get("/info", bRoutes.getInfo)

	//env
	app.Get("/env", bRoutes.getEnv)
}
