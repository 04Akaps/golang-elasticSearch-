package repository

import (
	"context"
	"elasticSearch/types"
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

func (s *Search) SearchUser(index string, query elastic.Query, size types.Size, text types.Sort) error {
	client := s.Client

	base := client.Search(index).Query(query).Pretty(true)

	if size.Size != 0 {
		base.Size(int(size.Size))
	}

	if len(text.Text) != 0 {
		for _, t := range text.Text {
			base.Sort(t, true)
		}
	}

	if err := checkIndexExisted(client, index); err != nil {
		return err
	} else if result, err := base.Do(context.TODO()); err != nil {
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
