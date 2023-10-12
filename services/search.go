package services

import (
	"elasticSearch/repository"
	"elasticSearch/types"
	"elasticSearch/types/schema"
	"github.com/olivere/elastic/v7"
)

//query.Must()      // -> AND 조건으로
//query.Filter()	// -> 점수에 영향을 안주고 필터를 한다?
//query.MustNot()   // -> Not 조건으로 동작
//query.Should()	// -> OR 조건으로 동작

//query.Boost()  	// -> 가중치를 부여
//query.MinimumNumberShouldMatch() // -> Should에서 최소한 만족해야 하는 쿼리의 수
//query.MinimumShouldMatch()		// -> 위에 있는 메서도 보다 좀 더 유연하게 만족해야 하는 수를 지정 ex : 백분율
//query.QueryName()					// 쿼리의 용도를 의미 -> 디버깅용으로 사용

/*
	Filter와 Must의 차이
	기본 동작은 동일, 특정 조건에 맞는 결과값을 보고 싶을 떄 사용
	하지만 Filter는 진짜 결과를 필터링 하는데에 사용이 된다. 무슨 말이냐면 일반적으로 elasticSearch는 검색 결과에 대해서 점수를 매기게 된다.
	높은 점수 일 수록, 쿼리의 결과에 적합하며 상위 포지션에 위치하게 된다.
		--> 일반적인 Must쿼리라면 모두 같은 점수를 가지게 될 것이다. : 어차피 해당 조건에 완전히 성립하는 데이터만 나오니깐

	여기서 Must가 이러한 역할을 수행하여 점수를 메기는 작업까지 진행이 되지만
	Filter는 검색에 대해서 점수를 메기지 않는다.
	그래서 쿼리에 대해서 더 빠르게 동작 할 수 있아.
	--> 점수를 메기지 않는 이유는 딱히 모른다.. 그냥 완전 거르는 용도로 사용을 할 떄만 사용을 하면 된다.
*/

type Search struct {
	search *repository.Search
}

func newSearchService(elasticSearch *repository.Search) *Search {
	return &Search{search: elasticSearch}
}

func (s *Search) FindAll(index string, size types.Size, text types.Sort) ([]*schema.Migration, error) {
	query := elastic.NewBoolQuery()
	return s.search.SearchMigration(index, query, size, text)
}

func (s *Search) SearchByName(index, name string, size types.Size, text types.Sort) ([]*schema.User, error) {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewMatchQuery("name", name))

	//elastic.NewMatchPhraseQuery("name", name) // 입력한 텍스트가 정확하게 연속이 되어야만 탐색이 이루어 진다.

	//query.Must(elastic.NewMatchQuery("name", name).Boost(2.0), elastic.NewMatchQuery("address", "c space e")) // 가중치를 준다. 해당 조건을 만족하는 결과를 더 상위에 위치 시킨다.
	//query.Must(elastic.NewMatchQuery("name", name), elastic.NewMatchQuery("address", "c space e")) // 이 경우 두 조건을 모두 성립해야 성공
	//query.Should(elastic.NewMatchQuery("name", name), elastic.NewTermQuery("address", "space")) // -> Or 조건으로 동작해서, 두 조건에 중 하나만 성립해도 가져 온다.

	//query.Should(elastic.NewTermQuery("address", "space")) // 이렇게 전송하면, 내 쿼리에 대해서 점수를 볼 수 있다.

	// elastic.NewWildcardQuery("address", "*space*") // 와일드 카드에 대한 쿼리를 전송 하는 방법

	return s.search.SearchUser(index, query, size, text)
}

func (s *Search) SearchByAddress(index, address string, size types.Size, text types.Sort) ([]*schema.User, error) {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewMatchQuery("address", address))
	return s.search.SearchUser(index, nil, size, text)
}

func (s *Search) SearchByAge(index string, age int64, size types.Size, text types.Sort) ([]*schema.User, error) {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewMatchQuery("age", age))

	//query.Filter(elastic.NewRangeQuery("age").Lt(100)) // 100미만의 값을 탐색 할 떄
	//query.Filter(elastic.NewRangeQuery("age").Lte(100)) // 100이하의 값을 탐색 할 떄
	//query.Filter(elastic.NewRangeQuery("age").Gt(100)) // 100 초과의 값을 탐색 할 떄
	//query.Filter(elastic.NewRangeQuery("age").Gte(100)) // 100 이상의 값을 탐색 할 떄

	return s.search.SearchUser(index, nil, size, text)
}
