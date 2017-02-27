package controllers

import (
	"encoding/json"
	"errors"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"github.com/dr013/product/app/models"
)

type ProductController struct {
	*revel.Controller
}

func (c ProductController) Index() revel.Result {
	var (
		products []models.Product
		err      error
	)
	products, err = models.GetProducts()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	c.Response.Status = 200
	return c.RenderJson(products)
}

func (c ProductController) Show(id string) revel.Result {
	var (
		product   models.Product
		err       error
		productID bson.ObjectId
	)

	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid product id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	productID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid product id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	product, err = models.GetProduct(productID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}

	c.Response.Status = 200
	return c.RenderJson(product)
}

func (c ProductController) Create() revel.Result {
	var (
		product models.Product
		err     error
	)

	err = json.NewDecoder(c.Request.Body).Decode(&product)
	if err != nil {
		errResp := buildErrResponse(err, "403")
		c.Response.Status = 403
		return c.RenderJson(errResp)
	}

	product, err = models.AddProduct(product)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	c.Response.Status = 201
	return c.RenderJson(product)
}

func (c ProductController) Update() revel.Result {
	var (
		product models.Product
		err     error
	)
	err = json.NewDecoder(c.Request.Body).Decode(&product)
	if err != nil {
		errResp := buildErrResponse(err, "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	err = product.UpdateProduct()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	return c.RenderJson(product)
}

func (c ProductController) Delete(id string) revel.Result {
	var (
		err       error
		product   models.Product
		productID bson.ObjectId
	)
	if id == "" {
		errResp := buildErrResponse(errors.New("Invalid product id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	productID, err = convertToObjectIdHex(id)
	if err != nil {
		errResp := buildErrResponse(errors.New("Invalid product id format"), "400")
		c.Response.Status = 400
		return c.RenderJson(errResp)
	}

	product, err = models.GetProduct(productID)
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	err = product.DeleteProduct()
	if err != nil {
		errResp := buildErrResponse(err, "500")
		c.Response.Status = 500
		return c.RenderJson(errResp)
	}
	c.Response.Status = 204
	return c.RenderJson(nil)
}
