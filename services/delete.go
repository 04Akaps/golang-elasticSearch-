package services

import (
	"elasticSearch/repository"
)

type Delete struct {
	elasticSearch *repository.Elastic
}

func newDeleteService(elasticSearch *repository.Elastic) *Delete {
	return &Delete{elasticSearch: elasticSearch}
}

func (d *Delete) DeleteUser(index, name string) error {
	return d.elasticSearch.Delete.DeleteUser(index, name)
}
