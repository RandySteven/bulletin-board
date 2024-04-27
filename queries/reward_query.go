package queries

const (
	InsertIntoReward GoQuery = `
		INSERT INTO rewards (name, description, image, user_id) 
		VALUES 
		    ($1, $2, $3, $4)
		RETURNING id
	`

	SelectAllRewards = `
		SELECT id, name, description, image, user_id, created_at, updated_at, deleted_at
		FROM rewards
	`

	SelectRewardByID = `
		SELECT id, name, description, image, user_id, created_at, updated_at, deleted_at
		FROM rewards
		WHERE id = $1
	`
)
