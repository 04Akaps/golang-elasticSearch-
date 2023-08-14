package repository

import (
	"context"
	"elasticSearch/types/schema"
	"errors"
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
	if user, ok := data.(schema.User); ok {
		client := c.Client
		if err := checkIndexExisted(client, index); err != nil {
			return err
		} else if _, err = client.Index().Index(index).Id(user.Name).BodyJson(data).Do(context.TODO()); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return errors.New("Not User Type")
	}
}
