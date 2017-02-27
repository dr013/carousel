package tests

import "github.com/revel/revel/testing"

type ProductTest struct {
        testing.TestSuite
}

func (t *ProductTest) Before() {
        println("Set up")
}

func (t *ProductTest) TestThatIndexPageWorks() {
        t.Get("/products")
        t.AssertOk()
        t.AssertContentType("application/json; charset=utf-8")
}

func (t *ProductTest) After() {
        println("Tear down")
}

