package service

import (
	"context"

	cs "github.com/JasurbekUz/catalog-service/genproto/catalog_service"
)

func (s *CatalogService) CreateBookCategory(ctx context.Context, req *cs.BookCategory) (*cs.Get_Book, error) {

	return &cs.Get_Book{}, nil
}

func (s *CatalogService) GetBookCategoryList(ctx context.Context, req *cs.ListReq) (*cs.BookCategoryListResp, error) {

	return &cs.BookCategoryListResp{}, nil
}

func (s *CatalogService) UpdateBookCategory(ctx context.Context, req *cs.BookCategory) (*cs.Get_Book, error) {

	return &cs.Get_Book{}, nil
}