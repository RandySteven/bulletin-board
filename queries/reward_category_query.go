package queries

const (
	InsertIntoRewardCategory GoQuery = `
		INSERT INTO reward_categories (reward_id, category_id)
		VALUES ($1, $2) RETURNING id
	`
)
