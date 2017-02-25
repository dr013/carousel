package controllers

import (
	"github.com/dr013/carousel/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var (
		schemas []models.Schema
		err     error
	)
	schemas, err = models.GetSchemas()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	c.Response.Status = 200
	return c.Render(schemas)
}
