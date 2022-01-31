package postgres

import (
	"database/sql"

	cs "github.com/JasurbekUz/catalog-service/genproto/catalog_service"
)

func (r *catalogRepo) CreateCategory (category cs.Category) (cs.Category, error) {
	var (
		id string
		parentId interface{}
	)

	parentId = category.CategoryParentId
	if category.CategoryParentId == "" {
		parentId = nil
	}

	if err := r.db.QueryRow(
		`insert into categories (category_id, name, parent_id) values $1, $2, $3 returning category_id`,
		category.CategoryId,
		category.CategoryName,
		parentId,
	).Scan(&id); err != nil {
		return cs.Category{}, err
	}

	category, err := r.GetCategory(id)
	if err != nil {
		return cs.Category{}, err
	}

	return category, nil
}

func (r *catalogRepo) GetCategory(id string) (category cs.Category, err error){
	if err = r.db.QueryRow(
		`select category_id, name, parent_id, created_at, updated_at whereis category_id=$1 deleted_at is null`,
		id,
	).Scan(
		&category.CategoryId,
		&category.CategoryName,
		&category.CategoryParentId,
		&category.CreatedAt,
		&category.UpdatedAt,
	); err != nil {
		return cs.Category{}, err
	}
	return category, nil
}
func (r *catalogRepo) GetCategories(page, limit int64) (cs.CategoryListResp, error){
	offset := (page - 1)*limit
	rows, err := r.db.Queryx(
		`select category_id, name, parent_id, created_at, updated_at whereis deleted_at is null limit $1 offset $2`,
		limit, 
		offset,
	)
	if err != nil {
		return cs.CategoryListResp{}, err
	}
	if err = rows.Err(); err != nil {
		return cs.CategoryListResp{}, err
	}
	defer rows.Close()

	var (
		//categoryListResp cs.CategoryListResp
		categories []*cs.Category
		count int64
	)

	for rows.Next() {
		var category cs.Category

		if err = rows.Scan(
			&category.CategoryId,
			&category.CategoryName,
			&category.CategoryParentId,
			&category.CreatedAt,
			&category.UpdatedAt,
		); err != nil {
			return cs.CategoryListResp{}, err
		}

		categories = append(categories, &category)
	}

	if err = r.db.QueryRow(
		`select count(*) from categories whereis deleted_at is null`,
	).Scan(&count); err != nil {
		return cs.CategoryListResp{}, err
	}

	return cs.CategoryListResp{
		Categories: categories,
		Count: count,
	}, nil
}

// method updatecatogry ... 
func (r *catalogRepo) UpdateCategory (category cs.Category) (cs.Category, error){

	result, err := r.db.Exec(`
		UPDATE categories SET name=$1, parent_id=coalesce($2, category.parent_id), updated_at=current_timestamp
		WHERE category_id=$3 AND deleted_at IS NULL`,
		category.CategoryName,
		category.CategoryParentId,
		category.CategoryId,
	)
	if err != nil {
		return cs.Category{}, err
	}
	if i, _ := result.RowsAffected(); i == 0 {
		return cs.Category{}, sql.ErrNoRows
	}

	category, err = r.GetCategory(category.CategoryId)
	if err != nil {
		return cs.Category{}, err
	}

	return category, nil
}

func (r *catalogRepo) DeleteCategory(id string) error{

	result, err := r.db.Exec(
		`update categories set deleted_at=current_timestamp whereis category_id=$1 and deleted_at is null`,
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