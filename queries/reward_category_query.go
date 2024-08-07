package queries

const (
	InsertIntoRewardCategory GoQuery = `
		INSERT INTO reward_categories (reward_id, category_id)
		VALUES ($1, $2) RETURNING id
	`

	SelectFromRewardCategory GoQuery = `
		SELECT id, reward_id, category_id, created_at, updated_at, deleted_at 
		FROM reward_categories
		AND deleted_at IS NULL
	`

	SelectByRewardID GoQuery = `
		SELECT id, reward_id, category_id, created_at, updated_at, deleted_at
		FROM reward_categories
		WHERE reward_id = $1
		AND deleted_at IS NULL
	`

	SelectByCategoryID GoQuery = `
		SELECT id, reward_id, category_id, created_at, updated_at, deleted_at
		FROM reward_categories
		WHERE category_id = $1
		AND deleted_at IS NULL
	`
)
