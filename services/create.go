package services

import (
	"elasticSearch/repository"
)

type Create struct {
	create *repository.Create
}

func newCreateService(create *repository.Create) *Create {
	return &Create{create: create}
}

func (c *Create) CreateData(index string, data interface{}) error {
	return c.create.CreateData(index, data)
}
