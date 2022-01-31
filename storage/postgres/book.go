package postgres

import (
	cs "github.com/JasurbekUz/catalog-service/genproto/catalog_service"
)

func (r *catalogRepo) CreateBook(cs.Book) (cs.Get_Book, error) {
	return cs.Get_Book{}, nil
}

func (r *catalogRepo) GetBook(id string) (cs.Get_Book, error) {
	return cs.Get_Book{}, nil
}

func (r *catalogRepo) GetBooks(page, limit int64) (cs.BookListResp, error) {
	return cs.BookListResp{}, nil
}

func (r *catalogRepo) UpdateBook(cs.Book) (cs.Get_Book, error) {
	return cs.Get_Book{}, nil
}

func (r *catalogRepo) DeleteBook(id string) error {
	return nil
}
