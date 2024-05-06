package queries

const (
	InsertCredit GoQuery = `
		INSERT INTO credits ('from_user_id', 'to_user_id', 'credit', 'description') 
		VALUES ($1, $2, $3, $4)
	`

	SelectUserCredit GoQuery = `
		SELECT id, from_user_id, to_user_id, credit, description, created_at, updated_at, deleted_at 
		FROM credits
		WHERE to_user_id = $1
	`
)
