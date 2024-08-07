package queries

const (
	SelectTaskRewardByTaskID GoQuery = `
		SELECT id, task_id, reward_id, created_at, updated_at, deleted_at
		FROM task_rewards
		WHERE task_id = $1
		AND deleted_at IS NULL
	`

	SelectTaskRewardByRewardID GoQuery = `
		SELECT id, task_id, reward_id, created_at, updated_at, deleted_at
		FROM task_rewards
		WHERE reward_id = $1
		AND deleted_at IS NULL
	`

	InsertIntoTaskReward = `
		INSERT INTO task_rewards(task_id, reward_id)
		VALUES 
		    ($1, $2) RETURNING id
	`
)
