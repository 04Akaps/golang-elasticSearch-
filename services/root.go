package services

import (
	"elasticSearch/config"
	"elasticSearch/repository"
	"github.com/inconshreveable/log15"
)

type ServiceRoot struct {
	elasticSearch *repository.Elastic
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

	if service.elasticSearch, err = repository.NewElastic(cfg); err != nil {
		return nil, err
	}

	service.Search = newSearchService(service.elasticSearch)
	service.Update = newUpdateService(service.elasticSearch)
	service.Delete = newDeleteService(service.elasticSearch)
	service.Create = newCreateService(service.elasticSearch)
	service.Admin = newAdminService(service.elasticSearch)

	return service, nil
}
