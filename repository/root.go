package repository

import (
	"elasticSearch/config"
	elsv8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/inconshreveable/log15"
	"github.com/olivere/elastic/v7"
)

type Elastic struct {
	Client   *elastic.Client
	V8Client *elsv8.Client

	logger log15.Logger

	Search *Search
	Create *Create
	Update *Update
	Delete *Delete
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

		elsV8Cfg := elsv8.Config{
			Addresses: []string{cfg.Elastic.Uri},
			Username:  cfg.Elastic.User,
			Password:  cfg.Elastic.Password,
		}

		if elasticClient.V8Client, err = elsv8.NewClient(elsV8Cfg); err != nil {
			return nil, err
		} else {

			elasticClient.Client = client

			elasticClient.Search = newSearch(client, elasticClient.V8Client)
			elasticClient.Create = newCreate(client)
			elasticClient.Admin = newAdmin(client)
			elasticClient.Update = newUpdate(client)
			elasticClient.Delete = newDelete(client)

			type ElsStatus struct {
				User     string `json:"user"`
				Password string `json:"password"`
			}
			status := &ElsStatus{
				User:     elasticCfg.User,
				Password: elasticCfg.Password,
			}

			elasticClient.logger.Info("Connected To ElasticSearch", "info", *status)

		}

		return elasticClient, nil
	}
}
