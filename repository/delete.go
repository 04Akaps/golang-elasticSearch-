package repository

import (
	"context"
	"github.com/olivere/elastic/v7"
)

type Delete struct {
	Client *elastic.Client
}

func newDelete(client *elastic.Client) *Delete {
	return &Delete{
		Client: client,
	}
}

func (d *Delete) DeleteUser(index, name string) error {
	client := d.Client
	if _, err := client.Delete().Index(index).Id(name).Do(context.TODO()); err != nil {
		return err
	} else {
		return nil
	}
}
