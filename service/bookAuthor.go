package service

import (
	"context"

	cs "github.com/JasurbekUz/catalog-service/genproto/catalog_service"
)


func (s *CatalogService) CreateBookAuthor(ctx context.Context, req *cs.BookAuthor) (*cs.Get_Book, error) {

	return &cs.Get_Book{}, nil
}

func (s *CatalogService) GetBookAuthorList(ctx context.Context, req *cs.ListReq) (*cs.BookAuthorListResp, error) {

	return &cs.BookAuthorListResp{}, nil
}

func (s *CatalogService) UpdateBookAuthor(ctx context.Context, req *cs.BookAuthor) (*cs.Get_Book, error) {

	return &cs.Get_Book{}, nil
}
