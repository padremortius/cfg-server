package v1

import (
	"cfg-server/internal/common"
	"cfg-server/internal/controller/structs"
	"cfg-server/internal/usecase/gitdata"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func getData(c *fiber.Ctx) error {
	app := strings.ToLower(c.Params("app", "test_app"))
	ext := strings.ToLower(c.Params("ext", "json"))

	data, err := gitdata.GetDataFromGit("", app)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.JSONResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	var raw []byte
	if ext == "yaml" || ext == "yml" {
		raw, _ = common.StructToYamlBytes(data)
		c.Response().Header.Set("Content-Type", "text/x-yaml")
	} else {
		c.Response().Header.Set("Content-Type", "application/json")
		raw, _ = common.StructToJSONBytes(data)
	}
	return c.Status(http.StatusOK).Send(raw)
}

func getDataWithEnv(c *fiber.Ctx) error {
	env := strings.ToLower(c.Params("env", ""))
	app := strings.ToLower(c.Params("app", "test_app"))
	ext := strings.ToLower(c.Params("ext", "json"))

	data, err := gitdata.GetDataFromGit(env, app)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.JSONResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	var raw []byte
	if ext == "yaml" || ext == "yml" {
		raw, _ = common.StructToYamlBytes(data)
		c.Response().Header.Set("Content-Type", "text/x-yaml")
	} else {
		c.Response().Header.Set("Content-Type", "application/json")
		raw, _ = common.StructToJSONBytes(data)
	}
	return c.Status(http.StatusOK).Send(raw)
}
