package repository

import (
	"context"
	"github.com/olivere/elastic/v7"
)

type Create struct {
	Client *elastic.Client
}

func newCreate(client *elastic.Client) *Create {
	return &Create{
		Client: client,
	}
}

func (c *Create) CreateData(index string, data interface{}) error {
	client := c.Client

	if err := checkIndexExisted(client, index); err != nil {
		return err
	} else if _, err = client.Index().Index(index).BodyJson(data).Do(context.TODO()); err != nil {
		return err
	} else {
		return nil
	}
}
