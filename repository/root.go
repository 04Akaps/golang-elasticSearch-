package repository

import (
	"elasticSearch/config"
	"github.com/inconshreveable/log15"
	"github.com/olivere/elastic/v7"
)

type Elastic struct {
	Client *elastic.Client
	logger log15.Logger

	Search *Search
	Create *Create
	Admin  *Admin
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

		elasticClient.Search = newSearch(client)
		elasticClient.Create = newCreate(client)
		elasticClient.Admin = newAdmin(client)

		type ElsStatus struct {
			User     string `json:"user"`
			Password string `json:"password"`
		}
		status := &ElsStatus{
			User:     elasticCfg.User,
			Password: elasticCfg.Password,
		}

		elasticClient.logger.Info("Connected To ElasticSearch", "info", *status)

		return elasticClient, nil
	}
}
