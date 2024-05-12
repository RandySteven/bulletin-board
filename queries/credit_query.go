package queries

const (
	InsertCredit GoQuery = `
		INSERT INTO credits (from_id, to_id, credit, description) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	SelectUserCredits GoQuery = `
		SELECT id, from_id, to_id, credit, description, created_at, updated_at, deleted_at 
		FROM credits
		WHERE to_id = $1
	`

	SelectCreditByID GoQuery = `
		SELECT id, from_id, to_id, credit, description, created_at, updated_at, deleted_at
		FROM credits
		WHERE id = $1
	`
)
