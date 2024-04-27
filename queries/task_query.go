package queries

const (
	InsertTask GoQuery = `INSERT INTO tasks (title, description, image, user_id, expired_date, status) 
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id
	`
	SelectAllTasks = `
		SELECT id, title, description, image, user_id, expired_date, created_at, updated_at, deleted_at
		FROM tasks
	`

	SelectTaskByID = `
		SELECT id, title, description, image, user_id, expired_date, created_at, updated_at, deleted_at
		FROM tasks
		WHERE id = $1
	`
)
