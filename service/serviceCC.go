package service

import (
	"github.com/JasurbekUz/catalog-service/pkg/logger"
	"github.com/JasurbekUz/catalog-service/storage"
)

type CatalogService struct {
	storage storage.Istorage
	logger logger.Logger
}

func NewCatalogService (storage storage.Istorage, log logger.Logger) *CatalogService {
	return &CatalogService{
		storage: storage,
		logger: log,
	}
}