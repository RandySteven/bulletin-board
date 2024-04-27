package queries

const (
	InsertUserTask GoQuery = `
		INSERT INTO user_tasks (user_id, task_id) 
		VALUES ($1, $2) RETURNING id
	`
)
