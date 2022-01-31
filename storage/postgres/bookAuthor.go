package postgres

import (

	cs "github.com/JasurbekUz/catalog-service/genproto/catalog_service"
)

func (r *catalogRepo) CreateBookAuthor(cs.BookAuthor) (cs.Get_Book, error) {

	return cs.Get_Book{}, nil
}

func (r *catalogRepo) GetBookAuthorList(page, limit int64) (cs.BookAuthorListResp, error) {

	return cs.BookAuthorListResp{}, nil
}

func (r *catalogRepo) UpdateBookAuthor(cs.BookAuthor) (cs.Get_Book, error) {

	return cs.Get_Book{}, nil
}