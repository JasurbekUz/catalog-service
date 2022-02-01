package service

import (
	"context"

	cs "github.com/JasurbekUz/catalog-service/genproto/catalog_service"
)

func (s *CatalogService) CreateBook(ctx context.Context, req *cs.Book) (*cs.Get_Book, error) {

	return &cs.Get_Book{} , nil
}

func (s *CatalogService) GetBook(ctx context.Context, req *cs.ByIdReq) (*cs.Get_Book, error) {
	
	return &cs.Get_Book{} , nil
}

func (s *CatalogService) GetBooks(ctx context.Context, req *cs.ListReq) (*cs.BookListResp, error) {
	
	return &cs.BookListResp{} , nil
}

func (s *CatalogService) UpdateBook(ctx context.Context, req *cs.Book) (*cs.Get_Book, error) {
	
	return &cs.Get_Book{} , nil
}

func (s *CatalogService) DeleteBook(ctx context.Context, req *cs.ByIdReq) (*cs.Empty, error) {
	
	return &cs.Empty{} , nil
}
