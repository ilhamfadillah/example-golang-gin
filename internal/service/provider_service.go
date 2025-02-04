package service

import (
	"example/apps/internal/model"
	"example/apps/internal/repository"
)

type ProviderService interface {
	GetAll() ([]*model.Provider, error)
	GetByUuid(uuid string) (*model.Provider, error)
	Store(provider *model.Provider) (*model.Provider, error)
	Update(uuid string, provider *model.Provider) (*model.Provider, error)
}

type providerService struct {
	repo repository.ProviderRepository
}

func NewProviderService(repo repository.ProviderRepository) ProviderService {
	return &providerService{repo: repo}
}

func (service *providerService) GetAll() ([]*model.Provider, error) {
	return service.repo.GetAll()
}

func (service *providerService) GetByUuid(uuid string) (*model.Provider, error) {
	return service.repo.GetByUuid(uuid)
}

func (service *providerService) Store(provider *model.Provider) (*model.Provider, error) {
	return service.repo.Store(provider)
}

func (service *providerService) Update(uuid string, provider *model.Provider) (*model.Provider, error) {
	return service.repo.Update(uuid, provider)
}
