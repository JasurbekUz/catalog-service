package repo

import (
	cs "github.com/JasurbekUz/catalog-service/genproto/catalog_service"
)

type CatalogStorageI interface {
	CreateAuthor(cs.Author) (cs.Author, error)
	GetAuthor(id string) (cs.Author, error)
	GetAuthors(page, limit int64) ([]*cs.Author, int64, error)
	UpdateAuthor(cs.Author) (cs.Author, error)
	DeleteAuthor(id string) error
}
