package queries

const (
	InsertIntoCategory  GoQuery = `INSERT INTO categories (category) VALUES ($1) RETURNING id`
	SelectCategoryByID          = `SELECT id, category, created_at, updated_at, deleted_at FROM categories WHERE id = $1`
	SelectAllCategories         = `SELECT id, category, created_at, updated_at, deleted_at FROM categories`
)
