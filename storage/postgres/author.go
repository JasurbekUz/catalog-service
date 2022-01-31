package postgres

import (
	"database/sql"

	cs "github.com/JasurbekUz/catalog-service/genproto/catalog_service"
)

func (r *catalogRepo) CreateAuthor(author cs.Author) (cs.Author, error) {
	var id string
	err := r.db.QueryRow(
		`insert into authors(author_id, name) values ($1, $2) returning author_id`,
		author.AuthorId,
		author.AuthorName,
	).Scan(&id)

	if err != nil {
		return cs.Author{}, err
	}

	author, err = r.GetAuthor(id)

	if err != nil {
		return cs.Author{}, err
	}

	return author, nil
}

func (r *catalogRepo) GetAuthor (id string) (author cs.Author, err error) {
	err = r.db.QueryRow(
		`select author_id, name, created_at, updated_at whereis author_id=&1 and deleted_at is null`,
		id,
	).Scan(
		&author.AuthorId,
		&author.AuthorName,
		&author.CreatedAt,
		&author.UpdatedAt,
	)

	if err != nil {
		return cs.Author{}, err
	}

	return author, nil
}

func (r *catalogRepo) GetAuthors(page, limit int64) ([]*cs.Author, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(
		`select author_id, name, created_at, updated_at whereis deleted_at is null limit $1 offset $2`,
		limit,
		offset,
	)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var (
		authors []*cs.Author
		count int64
	)

	for rows.Next() {
		var author cs.Author
		if err = rows.Scan(
			&author.AuthorId,
			&author.AuthorName,
			&author.CreatedAt,
			&author.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}
		authors = append(authors, &author)
	}

	if err = r.db.QueryRow(
		`select count(author_id) from authors where deleted_at is null`,
	).Scan(&count); err != nil {
		return nil, 0, err
	}
	return authors, count, nil
}

func (r *catalogRepo) UpdateAuthor(author cs.Author) (cs.Author, error) {
	result, err := r.db.Exec(
		`update authors set name=$1, updated_at=current_timestamp where author_id=$2 and deleted_at is null`,
		author.AuthorName,
		author.AuthorId,
	)
	if err != nil {
		return cs.Author{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return cs.Author{}, sql.ErrNoRows
	}

	author, err = r.GetAuthor(author.AuthorId)
	if err != nil{
		return cs.Author{}, err
	}

	return author, nil
}

func (r *catalogRepo) DeleteAuthor(id string) error {
	result, err := r.db.Exec(
		`update authors set deleted_at=current_timestamp where author_id=$1 and deleted_at is null`,
		id,
	)
	if err != nil {
		return err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}
	return nil
}
