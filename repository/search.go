package repository

import (
	"context"
	"elasticSearch/types"
	"elasticSearch/types/schema"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Search struct {
	Client *elastic.Client
}

func newSearch(client *elastic.Client) *Search {
	return &Search{
		Client: client,
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
