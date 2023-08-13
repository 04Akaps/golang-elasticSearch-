package services

import "elasticSearch/repository"

type Update struct {
	elasticSearch *repository.Elastic
}

func newUpdateService(elasticSearch *repository.Elastic) *Update {
	return &Update{elasticSearch: elasticSearch}
}
