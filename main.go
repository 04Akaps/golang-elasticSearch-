package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type Els struct {
	es *elastic.Client
}

func main() {
	if els, err := newEls(); err != nil {
		panic(err)
	} else {

		// els.insertData(INDEx, &Inner{
		// 	InnerName: "뮻",
		// })
		// els.createIndex(INDEx)

		rootBoolQuery := elastic.NewBoolQuery()
		els.searchData("collection-one", rootBoolQuery)
		// els.searchHighlightingData("test", rootBoolQuery)

		//query := elastic.NewTermQuery("name", "test1")
		//update := elastic.NewScript("ctx._source.inner.InnerName = params.name; ctx._source.age = params.age").Params(map[string]interface{}{
		//	"age":  500,
		//	"name": "dummy 2222",
		//})
		//
		//els.updateData("test", query, update)
	}
}

func newEls() (*Els, error) {
	if client, err := elastic.NewClient(
		elastic.SetBasicAuth(
			"hojin",
			"testHojin",
		),
		elastic.SetURL("http://localhost:9200/"),
		elastic.SetSniff(false),
	); err != nil {
		return nil, err
	} else {
		return &Els{es: client}, nil
	}
}

func (e *Els) insertData(index string, data interface{}) error {
	_, err := e.es.Index().
		Index(index).
		BodyJson(data).
		Do(context.Background())
	return err
}

func (e *Els) searchData(index string, query elastic.Query) error {
	if result, err := e.es.Search().
		Index(index).
		Query(query).
		Pretty(true).
		Size(100).
		// Sort("age", true).
		Do(context.Background()); err != nil {
		fmt.Println("여기서 EMa", err.Error())
		panic(err)
	} else {
		searchHit := result.Hits
		for _, v := range searchHit.Hits {
			fmt.Println(v)
		}
		return nil
	}
}

func (e *Els) searchHighlightingData(index string, query elastic.Query) {
	// 별 의미 x
	if searchResult, err := e.es.Search().
		Index(index).
		Query(query).
		Size(10).
		Sort("age", true).
		Highlight(
			elastic.NewHighlight().
				Field("name").     // 강조 표시할 필드 설정
				NumOfFragments(1). // 강조 표시할 단어 개수 설정
				PreTags("<b>").    // 강조 표시 시작 태그 설정
				PostTags("</b>"),  // 강조 표시 종료 태그 설정
		).
		Do(context.Background()); err != nil {
		panic(err)
	} else {
		searchHit := searchResult.Hits
		fmt.Println(searchHit)
		//for _, v := range searchHit.Hits {
		//	model := &DummyModel{}
		//	if err = json.Unmarshal(v.Source, model); err != nil {
		//		panic(err)
		//	}
		//	fmt.Println("name : ", model.Name, " Age : ", model.Age, " Address : ", model.Address, " Inner : ", model.Inner)
		//}
	}
}

func (e *Els) updateData(index string, query elastic.Query, update *elastic.Script) {
	_, err := e.es.UpdateByQuery(index).
		Query(query).
		Script(update).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
}

func (e *Els) DeleteData() {
}

func (e *Els) addAlias(indexName, aliasName string) error {
	_, err := e.es.Alias().Add(indexName, aliasName).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (e *Els) createIndex(indexName string) error {
	_, err := e.es.CreateIndex(indexName).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (e *Els) viewAllIndexes() error {
	if indexes, err := e.es.IndexNames(); err != nil {
		return err
	} else {
		for _, index := range indexes {
			fmt.Println(index)
		}
		return nil
	}
}
