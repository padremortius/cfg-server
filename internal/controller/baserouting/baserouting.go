package baserouting

import (
	"github.com/padremortius/cfg-server/internal/config"
	"github.com/padremortius/cfg-server/internal/controller/structs"

	fiber "github.com/gofiber/fiber/v2"
)

type (
	BaseRoutes struct {
		cfg config.Config
	}
)

func (b *BaseRoutes) getHealth(c *fiber.Ctx) error {
	return c.JSON(structs.Health{Status: "up"})
}

func (b *BaseRoutes) getInfo(c *fiber.Ctx) error {
	return c.JSON(b.cfg.Version)
}

func (b *BaseRoutes) getEnv(c *fiber.Ctx) error {
	return c.JSON(b.cfg)
}

func InitBaseRouter(app *fiber.App, aConfig config.Config) {
	bRoutes := BaseRoutes{cfg: aConfig}

	// K8s probe
	app.Get("/health", bRoutes.getHealth)

	//info about service
	app.Get("/info", bRoutes.getInfo)

	//env
	app.Get("/env", bRoutes.getEnv)
}
