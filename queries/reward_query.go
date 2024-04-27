package queries

const (
	InsertIntoReward GoQuery = `
		INSERT INTO rewards (name, description, image, user_id) 
		VALUES 
		    ($1, $2, $3, $4)
		RETURNING id
	`
)
