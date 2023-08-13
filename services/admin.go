package services

import (
	"context"
	"elasticSearch/repository"
)

type Admin struct {
	elasticSearch *repository.Elastic
}

func newAdminService(elasticSearch *repository.Elastic) *Admin {
	return &Admin{elasticSearch: elasticSearch}
}

func (a *Admin) CreateIndex(index string) error {
	instance := a.elasticSearch
	client := instance.Client

	if err := instance.CheckIndexExisted(index); err == nil {
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
	client := a.elasticSearch.Client
	if indexes, err := client.IndexNames(); err != nil {
		return nil, err
	} else {
		return indexes, nil
	}
}
