package repository

import (
	"context"
	"elasticSearch/types"
	"elasticSearch/types/schema"
	"encoding/json"
	"fmt"
	elsv8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/olivere/elastic/v7"
	"io"
)

type Search struct {
	Client   *elastic.Client
	V8Client *elsv8.Client
}

func newSearch(client *elastic.Client, V8Client *elsv8.Client) *Search {
	return &Search{
		Client:   client,
		V8Client: V8Client,
	}
}

func parsingV8Response(response *esapi.Response) ([]schema.HitsTypes, error) {
	defer response.Body.Close()

	if bytes, err := io.ReadAll(response.Body); err != nil {
		return nil, err
	} else {
		var responseType schema.V8ResponseType
		if err = json.Unmarshal(bytes, &responseType); err != nil {
			return nil, err
		} else {
			return responseType.Hits.Hits, nil
		}
	}
}

func (s *Search) SearchByV8Migration(index string, query map[string]interface{}) ([]*schema.Migration, error) {
	v8Client := s.V8Client

	if searchResponse, err := v8Client.Search(
		v8Client.Search.WithIndex(index),
		v8Client.Search.WithBody(esutil.NewJSONReader(query)),
	); err != nil || searchResponse.IsError() {
		return nil, err
	} else {

		if hits, err := parsingV8Response(searchResponse); err != nil {
			return nil, err
		} else {
			var migrationResponse []*schema.Migration
			for _, source := range hits {

				if responseByte, err := json.Marshal(source.Source); err != nil {
					continue
				} else {
					var m schema.Migration
					
					if err = json.Unmarshal(responseByte, &m); err != nil {
						continue
					} else {
						migrationResponse = append(migrationResponse, &m)
					}
				}

				// -> 타입이 map[string]interface이기 떄문에 reflect를 통한 type 체크 불필요
				//fmt.Println(source.Source)
				//fmt.Println(reflect.TypeOf(source.Source))
				//if reflect.TypeOf(source.Source) == reflect.TypeOf(schema.Migration{}) {
				//	migrationResponse = append(migrationResponse, source.Source.(*schema.Migration))
				//}
			}
			return migrationResponse, nil
		}
	}
}

func (s *Search) SearchMigration(index string, query elastic.Query, size types.Size, text types.Sort) ([]*schema.Migration, error) {
	client := s.Client

	base := client.Search(index).Query(query).Pretty(true)

	if len(text.Text) != 0 {
		for _, t := range text.Text {
			base.Sort(t, true)
		}
	}

	if size.From != 0 {
		base.From(int(size.From))
	}

	if size.Size != 0 {
		base.Size(int(size.Size))
	}

	if err := checkIndexExisted(client, index); err != nil {
		return nil, err
	} else if result, err := base.Do(context.TODO()); err != nil {
		return nil, err
	} else {
		var response []*schema.Migration
		searchHit := result.Hits
		for _, v := range searchHit.Hits {
			var model schema.Migration
			if err = json.Unmarshal(v.Source, &model); err != nil {
				return nil, err
			}

			response = append(response, &model)
		}

		return response, nil
	}
}

func (s *Search) SearchUser(index string, query elastic.Query, size types.Size, text types.Sort) ([]*schema.User, error) {
	client := s.Client

	base := client.Search(index).Query(query).Pretty(true)

	if len(text.Text) != 0 {
		for _, t := range text.Text {
			base.Sort(t, true)
		}
	}

	if size.From != 0 {
		base.From(int(size.From))
	}

	if size.Size != 0 {
		base.Size(int(size.Size))
	}

	if err := checkIndexExisted(client, index); err != nil {
		return nil, err
	} else if result, err := base.Do(context.TODO()); err != nil {
		return nil, err
	} else {
		var response []*schema.User
		searchHit := result.Hits
		for _, v := range searchHit.Hits {
			model := &schema.User{}
			if err = json.Unmarshal(v.Source, model); err != nil {
				return nil, err
			}
			fmt.Println(*v.Score)
			response = append(response, model)
		}

		return response, nil
	}
}
