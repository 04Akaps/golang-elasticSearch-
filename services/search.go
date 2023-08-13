package services

import (
	"elasticSearch/repository"
	"elasticSearch/types"
	"github.com/olivere/elastic/v7"
)

type Search struct {
	search *repository.Search
}

func newSearchService(elasticSearch *repository.Search) *Search {
	return &Search{search: elasticSearch}
}

func (s *Search) SearchByName(index, name string, size types.Size, text types.Sort) error {
	// TODO Query 작성
	query := elastic.NewBoolQuery()
	//query := elastic.NewTermQuery("name", name)
	return s.search.SearchUser(index, query, size, text)
}

func (s *Search) SearchByAddress(index, address string, size types.Size, text types.Sort) error {
	// TODO Query 작성
	return s.search.SearchUser(index, nil, size, text)
}

func (s *Search) SearchByAge(index string, age int64, size types.Size, text types.Sort) error {
	// TODO Query 작성
	return s.search.SearchUser(index, nil, size, text)
}

//
//rootBoolQuery := elastic.NewBoolQuery()
//els.searchData("collection-one", rootBoolQuery)
//// els.searchHighlightingData("test", rootBoolQuery)
//
////query := elastic.NewTermQuery("name", "test1")
////update := elastic.NewScript("ctx._source.inner.InnerName = params.name; ctx._source.age = params.age").Params(map[string]interface{}{
//// "age":  500,
//// "name": "dummy 2222",
////})
////
////els.updateData("test", query, update)
