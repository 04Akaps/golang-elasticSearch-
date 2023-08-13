package services

import (
	"context"
	"elasticSearch/repository"
)

type Create struct {
	elasticSearch *repository.Elastic
}

func newCreateService(elasticSearch *repository.Elastic) *Create {
	return &Create{elasticSearch: elasticSearch}
}

func (c *Create) CreateData(index string, data interface{}) error {
	instance := c.elasticSearch
	client := instance.Client

	if err := instance.CheckIndexExisted(index); err != nil {
		return err
	} else if _, err = client.Index().Index(index).BodyJson(data).Do(context.TODO()); err != nil {
		return err
	} else {
		return nil
	}
}
