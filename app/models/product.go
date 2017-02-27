package models

import (
        "gopkg.in/mgo.v2/bson"
        "time"
        "github.com/dr013/product/app/models/mongodb"
)

type Product struct {
        ID        bson.ObjectId `json:"id" bson:"_id"`
        Name      string        `json:"name" bson:"name"`
        Desc      string        `json:"desc" bson:"desc"`
        Jira      string        `json:"jira" bson:"jira"`
        Weight    int           `json:"weight" bson:"weight"`
        Project   string        `json:"project" bson:"project"`
        CreatedAt time.Time     `json:"created_at" bson:"created_at"`
        UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

func newProductCollection() *mongodb.Collection {
        return mongodb.NewCollectionSession("products")
}

// AddProduct insert a new Product into database and returns
// last inserted product on success.
func AddProduct(m Product) (product Product, err error) {
        c := newProductCollection()
        defer c.Close()
        m.ID = bson.NewObjectId()
        m.CreatedAt = time.Now()
        return m, c.Session.Insert(m)
}

// UpdateProduct update a Product into database and returns
// last nil on success.
func (m Product) UpdateProduct() error {
        c := newProductCollection()
        defer c.Close()

        err := c.Session.Update(bson.M{
                "_id": m.ID,
        }, bson.M{
                "$set": bson.M{
                        "name": m.Name, "desc": m.Desc, "jira": m.Jira, "project": m.Project, "weight": m.Weight, "updatedAt": time.Now()},
        })
        return err
}

// DeleteProduct Delete Product from database and returns
// last nil on success.
func (m Product) DeleteProduct() error {
        c := newProductCollection()
        defer c.Close()

        err := c.Session.Remove(bson.M{"_id": m.ID})
        return err
}

// GetProducts Get all Product from database and returns
// list of Product on success
func GetProducts() ([]Product, error) {
        var (
                products []Product
                err      error
        )

        c := newProductCollection()
        defer c.Close()

        err = c.Session.Find(nil).Sort("-createdAt").All(&products)
        return products, err
}

// GetProduct Get a Product from database and returns
// a Product on success
func GetProduct(id bson.ObjectId) (Product, error) {
        var (
                product Product
                err     error
        )

        c := newProductCollection()
        defer c.Close()

        err = c.Session.Find(bson.M{"_id": id}).One(&product)
        return product, err
}
