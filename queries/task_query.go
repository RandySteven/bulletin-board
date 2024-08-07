package queries

const (
	InsertTask GoQuery = `INSERT INTO tasks (title, description, image, user_id, expired_date, status) 
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id
	`
	SelectAllTasks = `
		SELECT id, title, description, image, status, user_id, expired_date, created_at, updated_at, deleted_at
		FROM tasks
		WHERE deleted_at IS NULL
	`

	SelectTaskByID = `
		SELECT id, title, description, image, status, user_id, expired_date, created_at, updated_at, deleted_at
		FROM tasks
		WHERE id = $1
		AND deleted_at IS NULL
	`

	SelectTasksThatAlreadyExpired = `
		SELECT * FROM tasks
		WHERE expired_date <= NOW()
		AND deleted_at IS NULL
	`

	UpdateTaskExpiryDate = `
		UPDATE tasks
			SET status = 'off', updated_at = NOW()
		WHERE expired_date < NOW() AND status = 'open'
		AND deleted_at IS NULL
	`
)
