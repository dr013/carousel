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
                err      error
        )
        products, err = models.GetProducts()
        schemas, err = models.GetSchemas()
        c.Session["testing"] = "1"
        c.Session["REVEL_LANG"] = c.Request.Locale
        if err != nil {
                errResp := buildErrResponse(err, "500")
                c.Response.Status = 500
                return c.RenderJson(errResp)
        }
        c.Response.Status = 200
        d := time.Now().Year()
        return c.Render(products, schemas, d)
}

func (c App) Login() revel.Result {
        c.Session["REVEL_LANG"] = c.Request.Locale
        user := c.Params.Form.Get("user")
        passwd := c.Params.Form.Get("passwd")
        revel.INFO.Println("Process login for ", user)
        if user == "krukov" {
                c.Session["auth"] = "true"
                c.Session["user"] = user
                c.Session["passwd"] = passwd
                revel.INFO.Println("Logged as ", user)
        }

        return c.Redirect(App.Index)
}

func (c App) Logout() revel.Result {
        revel.INFO.Println("Logout for ", c.Session["user"])
        delete(c.Session, "auth")
        delete(c.Session, "user")
        return c.Redirect(App.Index)
}
