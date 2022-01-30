package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/JasurbekUz/catalog-service/storage/postgres"
	"github.com/JasurbekUz/catalog-service/storage/repo"
)

// istorage
type Istorage interface {
	Catalog() repo.CatalogStorageI
}

type storagePg struct {
	db        *sqlx.DB
	CatalogRepo repo.CatalogStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:        db,
		CatalogRepo: postgres.NewCatalogRepo(db),
	}
}

func (s storagePg) Catalog() repo.CatalogStorageI {
	return s.CatalogRepo
}
