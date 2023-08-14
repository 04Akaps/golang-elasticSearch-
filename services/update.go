package services

import (
	"elasticSearch/repository"
	"github.com/olivere/elastic/v7"
)

type Update struct {
	elasticSearch *repository.Elastic
}

func newUpdateService(elasticSearch *repository.Elastic) *Update {
	return &Update{elasticSearch: elasticSearch}
}

func (u *Update) UpDateUser(index, name string, age int64) error {
	query := elastic.NewBoolQuery()
	// 정확한 name값으 쿼리
	query.Must(elastic.NewMatchQuery("name.keyword", name))
	// script로 짜줘야 한다.
	script := elastic.NewScript("ctx._source.age = params.age").Param("age", age)

	// 만약 여러개의 값을 업데이트 하려면
	//update := elastic.NewScript("ctx._source.inner.InnerName = params.name; ctx._source.age = params.age").Params(map[string]interface{}{
	//	"age":  500,
	//	"name": "dummy 2222",
	//})
	// 해당 코드는 inner이라는 값은 object로 구성이 되어 있고, 해당 값에서 InnerName값을 params로 지정하고, age값 또한 age로 지정하라는 의미

	return u.elasticSearch.Update.UpdateUser(index, query, script)
}
