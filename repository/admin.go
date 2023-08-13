package repository

import (
	"context"
	"github.com/olivere/elastic/v7"
)

type Admin struct {
	client *elastic.Client
}

func newAdmin(client *elastic.Client) *Admin {
	return &Admin{
		client: client,
	}
}

func (a *Admin) CreateIndex(index string) error {
	client := a.client

	if err := checkIndexExisted(client, index); err == nil {
		// 인덱스가 이미 존재 하는 경우
		return nil
	} else if _, err = client.CreateIndex(index).Do(context.TODO()); err != nil {
		// 없다면 생성
		return err
	} else {
		return nil
	}
}

func (a *Admin) ViewAllIndexes() ([]string, error) {
	client := a.client
	if indexes, err := client.IndexNames(); err != nil {
		return nil, err
	} else {
		return indexes, nil
	}
}
