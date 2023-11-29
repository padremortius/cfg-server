package v1

import (
	"cfg-server/internal/common"
	"cfg-server/internal/controller/structs"
	"cfg-server/internal/usecase/gitdata"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func getDataInJSON(c *fiber.Ctx) error {
	app := strings.ToLower(c.Params("app", "test_app"))

	data, err := gitdata.GetDataFromGit("", app)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.JSONResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(data)
}

func getDataInYaml(c *fiber.Ctx) error {
	app := strings.ToLower(c.Params("app", "test_app"))

	data, err := gitdata.GetDataFromGit("", app)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.JSONResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	raw, _ := common.StructToYamlBytes(data)
	return c.Status(http.StatusOK).Send(raw)
}

func getDataWithEnvInJSON(c *fiber.Ctx) error {
	app := strings.ToLower(c.Params("app", "test_app"))
	env := strings.ToLower(c.Params("env", ""))

	data, err := gitdata.GetDataFromGit(env, app)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.JSONResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.Status(http.StatusOK).JSON(data)
}

func getDataWithEnvInYaml(c *fiber.Ctx) error {
	app := strings.ToLower(c.Params("app", "test_app"))
	env := strings.ToLower(c.Params("env", ""))
	data, err := gitdata.GetDataFromGit(env, app)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(structs.JSONResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	raw, _ := common.StructToYamlBytes(data)
	return c.Status(http.StatusOK).Send(raw)
}
