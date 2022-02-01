package postgres

import (
	//"database/sql"

	//"github.com/huandu/go-sqlbuilder"
	
	cs "github.com/JasurbekUz/catalog-service/genproto/catalog_service"
)

func (r *catalogRepo) CreateBook(book cs.Book) (cs.Get_Book, error) {
	var id string
	if err := r.db.QueryRow(
		`inster into books (book_id, name) values ($1, $2) returning book_id`,
		book.BookId,
		book.BookName,
	).Scan(&id); err != nil {
		return cs.Get_Book{}, err
	}

	get_book, err := r.GetBook(id)
	if err != nil {
		return cs.Get_Book{}, err
	}

	return get_book, nil
}

func (r *catalogRepo) GetBook(id string) (book cs.Get_Book, err error) {
	var (
		authors []*cs.Author
		categories []*cs.Category
	)

	if err = r.db.QueryRow(
		`select book_id, name, created_at, updated_at from books whereis deleted_at is null`,
	).Scan(
		&book.BookId,
		&book.BookName,
	); err != nil {
		return cs.Get_Book{}, err
	}
	// cl func for get book authors
	authors, _ = func(id string) (auths []*cs.Author, err error) {
		rows, err := r.db.Queryx(
			`select a.author_id, a.name, a.created_at, a.updated_at from bookAuthor as b join authors as a using whereis b.book_id=$1 and a.deleted_at is null`,
			id,
		)
		if err != nil {
			return nil, err
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var author cs.Author
			if err = rows.Scan(
				&author.AuthorId,
				&author.AuthorName,
				&author.CreatedAt,
				&author.UpdatedAt,
			); err != nil {
				return nil, err
			}
			auths = append(auths, &author)
		}

		return auths, nil
	}(book.BookId)	

	categories, _ = func(id string) (categs []*cs.Category, err error) {
		rows, err := r.db.Queryx(
			`elect c.category_id, c.name, c.parent_id, c.created_at, c.updated_at from bookCategory as b join categories as c using whereis b.book_id=$1 and whereis c.deleted_at is null`,
			id,
		)
		if err != nil {
			return nil, err
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
		var category cs.Category

		if err = rows.Scan(
			&category.CategoryId,
			&category.CategoryName,
			&category.CategoryParentId,
			&category.CreatedAt,
			&category.UpdatedAt,
		); err != nil {
			return nil , err
		}

		categs = append(categs, &category)
	}

		return categs, nil
	}(book.BookId)

	return cs.Get_Book{
		BookId: book.BookId,
		BookName: book.BookName,
		Authors: authors,
		Categories: categories,
	}, nil
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
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*func (r *catalogRepo) GetBooks(page, limit int64, filters map[string]string) ([]*pb.BookResp, int64, error) {
	offset := (page - 1) * limit
	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("b.book_id", "b.name", "b.author_id", "b.created_at", "b.updated_at")
	sb.From("book b")
	sb.Where("b.deleted_at IS NULL")

	if val, ok := filters["category"]; ok && val != "" {
		args := utils.StringSliceToInterfaceSlice(utils.ParseFilter(val))
		sb.JoinWithOption("LEFT", "book_category bc", "b.book_id=bc.book_id")
		sb.Where(sb.In("bc.category_id", args...))
	}

	if val, ok := filters["author"]; ok && val != "" {
		args := utils.StringSliceToInterfaceSlice(utils.ParseFilter(val))
		sb.JoinWithOption("LEFT", "author a", "b.author_id=a.author_id")
		sb.Where(sb.In("b.author_id", args...))
	}

	sb.Where("(SELECT count(*) FROM author WHERE deleted_at IS NULL AND author_id=b.author_id) <> 0")
	sb.GroupBy("b.book_id", "b.name")
	sb.Limit(int(limit))
	sb.Offset(int(offset))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() // nolint:errcheck
	var (
		books []*pb.BookResp
		count int64
	)

	for rows.Next() {
		var (
			book     pb.BookResp
			authorId string
			author   pb.Author
		)
		err = rows.Scan(
			&book.Id,
			&book.Name,
			&authorId,
			&book.CreatedAt,
			&book.UpdatedAt,
		)

		if err != nil {
			return nil, 0, err
		}

		author, err = r.GetAuthor(authorId)
		if err != nil {
			return nil, 0, err
		}

		book.Author = &author
		books = append(books, &book)
	}

	for _, book := range books {
		var categories []*pb.Category
		categories, err = GetBookCategory(r, book.Id)
		if err != nil {
			return nil, 0, err
		}
		book.Category = categories
	}

	sbc := sqlbuilder.NewSelectBuilder()
	sbc.Select("ROW_NUMBER() over (order by b.book_id) as r_number")
	sbc.From("book b")
	sbc.Where("b.deleted_at IS NULL")

	if val, ok := filters["category"]; ok && val != "" {
		args := utils.StringSliceToInterfaceSlice(utils.ParseFilter(val))
		sbc.JoinWithOption("LEFT", "book_category bc", "b.book_id=bc.book_id")
		sbc.Where((sbc.In("bc.category_id", args...)))
	}

	if val, ok := filters["author"]; ok && val != "" {
		args := utils.StringSliceToInterfaceSlice(utils.ParseFilter(val))
		sbc.JoinWithOption("LEFT", "author a", "b.author_id=a.author_id")
		sbc.Where(sbc.In("b.author_id", args...))
	}

	sbc.GroupBy("b.book_id")
	sbc.OrderBy("r_number desc")
	sbc.Limit(1)
	query, args = sbc.BuildWithFlavor(sqlbuilder.PostgreSQL)
	err = r.db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return books, count, nil*/