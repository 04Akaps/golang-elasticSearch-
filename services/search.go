package services

import (
	"context"
	"elasticSearch/repository"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Search struct {
	elasticSearch *repository.Elastic
}

func newSearchService(elasticSearch *repository.Elastic) *Search {
	return &Search{elasticSearch: elasticSearch}
}

func (s *Search) SearchData(index string, query elastic.Query) error {
	instance := s.elasticSearch
	client := instance.Client

	if err := instance.CheckIndexExisted(index); err != nil {
		return err
	} else if result, err := client.Search(index).Query(query).Pretty(true).Size(100).Do(context.TODO()); err != nil {
		return err
	} else {
		searchHit := result.Hits
		for _, v := range searchHit.Hits {
			fmt.Println(v)
			//model := &DummyModel{}
			//	if err = json.Unmarshal(v.Source, model); err != nil {
			//		panic(err)
			//	}
			//	fmt.Println("name : ", model.Name, " Age : ", model.Age, " Address : ", model.Address, " Inner : ", model.Inner)
		}

		return nil
	}
}
