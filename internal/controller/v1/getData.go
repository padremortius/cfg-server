package v1

import (
	"cfg-server/internal/common"
	"cfg-server/internal/controller/structs"
	"cfg-server/internal/usecase/gitdata"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func getProfileAndName(aName string) (string, string) {
	items := strings.Split(aName, "-")
	count := len(items)
	if count == 1 {
		return aName, ""
	}
	appName := strings.Join(items[:count-1], "-")
	return appName, items[count-1]
}

func getData(c *fiber.Ctx) error {
	ext := strings.ToLower(c.Params("ext", "json"))
	app, profile := getProfileAndName(strings.ToLower(c.Params("app", "test_app")))

	data, err := gitdata.GetDataFromGit("", app, profile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.JSONResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	var raw []byte
	if ext == "yaml" || ext == "yml" {
		raw, _ = common.StructToYamlBytes(data)
		c.Response().Header.Set("Content-Type", "text/x-yaml; charset=UTF-8")
	} else {
		if ext == "json" {
			c.Response().Header.Set("Content-Type", "application/json")
			raw, _ = common.StructToJSONBytes(data)
		} else {
			return c.Status(http.StatusNotFound).JSON(structs.JSONResult{Code: http.StatusNotFound, Message: "Unknown format file"})
		}
	}
	return c.Status(http.StatusOK).Send(raw)
}

func getDataWithEnv(c *fiber.Ctx) error {
	env := strings.ToLower(c.Params("env", ""))
	app, profile := getProfileAndName(strings.ToLower(c.Params("app", "test_app")))
	ext := strings.ToLower(c.Params("ext", "json"))

	data, err := gitdata.GetDataFromGit(env, app, profile)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.JSONResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	var raw []byte
	switch ext {
	case "yaml", "yml":
		{
			raw, _ = common.StructToYamlBytes(data)
			c.Response().Header.Set("Content-Type", "text/x-yaml")
		}
	case "json":
		{
			c.Response().Header.Set("Content-Type", "application/json")
			raw, _ = common.StructToJSONBytes(data)
		}
	default:
		{
			c.Response().Header.Set("Content-Type", "text/plain")
			raw, _ = common.StructToTomlBytes(data)
		}
	}
	return c.Status(http.StatusOK).Send(raw)
}
