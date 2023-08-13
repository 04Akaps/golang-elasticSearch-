package services

import "elasticSearch/repository"

type Create struct {
	elasticSearch *repository.Elastic
}

func newCreateService(elasticSearch *repository.Elastic) *Create {
	return &Create{elasticSearch: elasticSearch}
}
