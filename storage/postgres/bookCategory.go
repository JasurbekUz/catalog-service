package postgres

import (

	cs "github.com/JasurbekUz/catalog-service/genproto/catalog_service"
)

func (r *catalogRepo) CreateBookCategory(cs.BookCategory) (cs.Get_Book, error) {
	
	return cs.Get_Book{}, nil
}

func (r *catalogRepo) GetBookCategoryList(page, limit int64) (cs.BookCategoryListResp, error) {
	
	return cs.BookCategoryListResp{}, nil
}

func (r *catalogRepo) UpdateBookCategory(cs.BookCategory) (cs.Get_Book, error) {
	
	return cs.Get_Book{}, nil
}