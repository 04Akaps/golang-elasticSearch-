package repository

import (
	"context"
	"github.com/olivere/elastic/v7"
)

type Update struct {
	Client *elastic.Client
}

func newUpdate(client *elastic.Client) *Update {
	return &Update{
		Client: client,
	}
}

func (c *Update) UpdateUser(index string, query elastic.Query, script *elastic.Script) error {
	client := c.Client

	if err := checkIndexExisted(client, index); err != nil {
		return err
	} else if _, err = client.UpdateByQuery(index).Query(query).Script(script).Do(context.TODO()); err != nil {
		return err
	} else {
		return nil
	}
}
