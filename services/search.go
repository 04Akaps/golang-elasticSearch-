package services

import (
	"elasticSearch/repository"
	"elasticSearch/types"
)

type Search struct {
	search *repository.Search
}

func newSearchService(elasticSearch *repository.Search) *Search {
	return &Search{search: elasticSearch}
}

func (s *Search) SearchByName(index, name string, size types.Size, text types.Sort) error {
	// TODO Query 작성
	//query := elastic.NewBoolQuery()
	//elastic.NewTermQuery()
	return s.search.SearchUser(index, nil, size, text)
}

func (s *Search) SearchByAddress(index, address string, size types.Size, text types.Sort) error {
	// TODO Query 작성
	return s.search.SearchUser(index, nil, size, text)
}

func (s *Search) SearchByAge(index string, age int64, size types.Size, text types.Sort) error {
	// TODO Query 작성
	return s.search.SearchUser(index, nil, size, text)
}
