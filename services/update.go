package services

import (
	"elasticSearch/repository"
)

type Update struct {
	elasticSearch *repository.Elastic
}

func newUpdateService(elasticSearch *repository.Elastic) *Update {
	return &Update{elasticSearch: elasticSearch}
}

//func (u *Update) UpDateData(index string, query elastic.Query, script *elastic.Script) error {
//	instance := u.elasticSearch
//	client := instance.Client
//
//	if err := instance.CheckIndexExisted(index); err != nil {
//		return err
//	} else if _, err = client.UpdateByQuery(index).Query(query).Script(script).Do(context.TODO()); err != nil {
//		return err
//	} else {
//
//		return nil
//	}
//}
