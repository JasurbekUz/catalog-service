package postgres

import (
	//"database/sql"

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
	return nil, 0, nil
}

func (r *catalogRepo) UpdateAuthor(cs.Author) (cs.Author, error) {
	return cs.Author{}, nil
}

func (r *catalogRepo) DeleteAuthor(id string) error {
	return nil
}
