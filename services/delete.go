package services

import "elasticSearch/repository"

type Delete struct {
	elasticSearch *repository.Elastic
}

func newDeleteService(elasticSearch *repository.Elastic) *Delete {
	return &Delete{elasticSearch: elasticSearch}
}
