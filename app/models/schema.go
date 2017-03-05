package models

import (
	"github.com/dr013/carousel/app/models/mongodb"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Schema struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Task       string        `json:"task" bson:"task"`
	DbName     string        `json:"db_name" bson:"db_name"`
	DbLogin    string        `json:"db_login" bson:"db_login"`
	DbHost     string        `json:"db_host" bson:"db_host"`
	DbPassword string        `json:"db_password" bson:"db_password"`
	DbPort     int           `json:"db_port" bson:"db_port"`
	DbSize     int           `json:"db_size" bson:"db_size"`
	LockedDate time.Time     `json:"locked_date" bson:"locked_date"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" bson:"updated_at"`
}

func newSchemaCollection() *mongodb.Collection {
	return mongodb.NewCollectionSession("schema_list")
}

// AddSchema insert a new Schema into database and returns
// last inserted schema on success.
func AddSchema(m Schema) (schema Schema, err error) {
	c := newSchemaCollection()
	defer c.Close()
	m.ID = bson.NewObjectId()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return m, c.Session.Insert(m)
}

// UpdateSchema update a Schema into database and returns
// last nil on success.
func (m Schema) UpdateSchema() error {
	c := newSchemaCollection()
	defer c.Close()

	err := c.Session.Update(bson.M{
		"_id": m.ID,
	}, bson.M{
		"$set": bson.M{
			"task": m.Task, "db_name": m.DbName, "db_login": m.DbLogin, "db_host": m.DbHost, "db_password": m.DbPassword, "db_port": m.DbPort, "db_size": m.DbSize, "locked_date": m.LockedDate, "updatedAt": time.Now()},
	})
	return err
}

// DeleteSchema Delete Schema from database and returns
// last nil on success.
func (m Schema) DeleteSchema() error {
	c := newSchemaCollection()
	defer c.Close()

	err := c.Session.Remove(bson.M{"_id": m.ID})
	return err
}

// GetSchemas Get all Schema from database and returns
// list of Schema on success
func GetSchemas() ([]Schema, error) {
	var (
		schemas []Schema
		err     error
	)

	c := newSchemaCollection()
	defer c.Close()

	err = c.Session.Find(nil).Sort("-createdAt").All(&schemas)
	return schemas, err
}

// GetSchema Get a Schema from database and returns
// a Schema on success
func GetSchema(id bson.ObjectId) (Schema, error) {
	var (
		schema Schema
		err    error
	)

	c := newSchemaCollection()
	defer c.Close()

	err = c.Session.Find(bson.M{"_id": id}).One(&schema)
	return schema, err
}

//func CheckToken(token string, project string) bool{
//	var (
//		result bool
//	        err error
//		instance &Project
//	)
//
//	c := NewProjectCollection()
//	defer c.Close()
//
//	err = c.Session.Find(bson.M{"project": project}).One(&instance)
//
//	if instance == nil {
//
//	}
//
//
//	return result
//}
