package controllers

import (
	"encoding/json"
	"errors"
	"github.com/dr013/carousel/app/models"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type SchemaController struct {
	*revel.Controller
}

func (c SchemaController) Index() revel.Result {
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
	return c.RenderJson(schemas)
}

func (c SchemaController) Show(id string) revel.Result {
	var (
		schema   models.Schema
		err      error
		schemaID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid schema id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	schemaID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid schema id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	schema, err = models.GetSchema(schemaID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}

	c.Response.Status = 200
	return c.RenderJson(schema)
}

func (c SchemaController) Create() revel.Result {
	var (
		schema models.Schema
		err    error
	)

	err = json.NewDecoder(c.Request.Body).Decode(&schema)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJson(errResp)
	}

	schema, err = models.AddSchema(schema)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	c.Response.Status = 201
	return c.RenderJson(schema)
}

func (c SchemaController) Update() revel.Result {
	var (
		schema models.Schema
		err    error
	)
	err = json.NewDecoder(c.Request.Body).Decode(&schema)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	err = schema.UpdateSchema()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	return c.RenderJson(schema)
}

func (c SchemaController) Delete(id string) revel.Result {
	var (
		err      error
		schema   models.Schema
		schemaID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid schema id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	schemaID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid schema id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	schema, err = models.GetSchema(schemaID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	err = schema.DeleteSchema()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	c.Response.Status = 204
	return c.RenderJson(nil)
}
