package services

import (
	"elasticSearch/config"
	"elasticSearch/repository"
	"github.com/inconshreveable/log15"
)

type ServiceRoot struct {
	ElasticSearch *repository.Elastic
	logger        log15.Logger

	Search *Search
	Update *Update
	Delete *Delete
	Create *Create
	Admin  *Admin
}

func NewService(cfg *config.Config) (*ServiceRoot, error) {
	service := &ServiceRoot{
		logger: log15.New("module", "service"),
	}
	var err error

	if service.ElasticSearch, err = repository.NewElastic(cfg); err != nil {
		return nil, err
	}

	service.Search = newSearchService(service.ElasticSearch.Search)
	service.Update = newUpdateService(service.ElasticSearch)
	service.Delete = newDeleteService(service.ElasticSearch)
	service.Create = newCreateService(service.ElasticSearch.Create)
	service.Admin = newAdminService(service.ElasticSearch.Admin)

	return service, nil
}
