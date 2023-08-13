package services

import "elasticSearch/repository"

type Search struct {
	elasticSearch *repository.Elastic
}

func newSearchService(elasticSearch *repository.Elastic) *Search {
	return &Search{elasticSearch: elasticSearch}
}
