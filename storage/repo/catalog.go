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

	CreateCategory(cs.Category) (cs.Category, error)
    GetCategory(id string) (cs.Category, error)
    GetCategories(page, limit int64) (cs.CategoryListResp, error)
    UpdateCategory(cs.Category) (cs.Category, error)
    DeleteCategory(id string) error

	CreateBook(cs.Book) (cs.Get_Book, error);
    GetBook(id string) (cs.Get_Book, error);
    GetBooks(page, limit int64) (cs.BookListResp, error);
    UpdateBook(cs.Book) (cs.Get_Book, error);
    DeleteBook(id string) error;

	CreateBookCategory(cs.BookCategory) (cs.Get_Book, error)
    GetBookCategoryList(page, limit int64) (cs.BookCategoryListResp, error)
    UpdateBookCategory(cs.BookCategory) (cs.Get_Book, error)

	CreateBookAuthor(cs.BookAuthor) (cs.Get_Book, error)
    GetBookAuthorList(page, limit int64) (cs.BookAuthorListResp, error)
    UpdateBookAuthor(cs.BookAuthor) (cs.Get_Book, error)
}
