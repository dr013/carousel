package controllers

import (
	"github.com/revel/revel"
	"github.com/dr013/carousel/app/models"
	"time"
)

type App struct {
	*revel.Controller
}


func (c App) Index() revel.Result {
	var (
		products []models.Product
		schemas  []models.Schema
		err     error
	)
	products, err = models.GetProducts()
	schemas, err = models.GetSchemas()

	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	c.Response.Status = 200
	d := time.Now().Year()
	return c.Render(products, schemas, d)
}
