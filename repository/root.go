package repository

import (
	"context"
	"elasticSearch/config"
	"errors"
	"fmt"
	"github.com/inconshreveable/log15"
	"github.com/olivere/elastic/v7"
)

type Elastic struct {
	Client *elastic.Client
	logger log15.Logger
}

func NewElastic(cfg *config.Config) (*Elastic, error) {
	elasticCfg := cfg.Elastic
	elasticClient := &Elastic{
		logger: log15.New("module", "repository/elastic"),
	}

	if client, err := elastic.NewClient(
		elastic.SetBasicAuth(
			elasticCfg.User,
			elasticCfg.Password,
		),
		elastic.SetURL(elasticCfg.Uri),
		elastic.SetSniff(false),
	); err != nil {
		return nil, err
	} else {
		elasticClient.Client = client
		return elasticClient, nil
	}
}

func (els *Elastic) CheckIndexExisted(index string) error {
	ctx := context.TODO()
	indices := []string{index}
	client := els.Client

	existService := elastic.NewIndicesExistsService(client)
	existService.Index(indices)

	exist, err := existService.Do(ctx)

	if err != nil {
		message := fmt.Sprintf("NewIndicesExistsService.Do() %s", err.Error())
		return errors.New(message)
	} else if !exist {
		fmt.Println("nOh no! The index", index, "doesn't exist.")
		fmt.Println("Create the index, and then run the Go script again")
		if _, err = client.CreateIndex(index).Do(ctx); err != nil {
			return err
		} else {
			return nil
		}
	} else if exist {
		fmt.Println("Index name:", index, " exists!")
		return nil
	} else {
		return nil
	}
}
