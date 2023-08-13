package services

import (
	"elasticSearch/repository"
)

type Admin struct {
	admin *repository.Admin
}

func newAdminService(admin *repository.Admin) *Admin {
	return &Admin{admin: admin}
}

func (a *Admin) CreateIndex(index string) error {
	return a.admin.CreateIndex(index)
}

func (a *Admin) ViewAllIndexes() ([]string, error) {
	return a.admin.ViewAllIndexes()
}
