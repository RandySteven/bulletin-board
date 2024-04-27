package queries

const (
	InsertIntoRelation GoQuery = `
		INSERT INTO relations (user_id, friend_id, status)
		VALUES 
		    ($1, $2, $3)
		RETURNING id
	`

	SelectUserRelations = `SELECT id, user_id, friend_id, status, created_at, updated_at, deleted_at
		FROM relations
		WHERE user_id = $1
	`

	SelectFriendRelation = `SELECT id, user_id, friend_id, status, created_at, updated_at, deleted_at
		FROM relations
		WHERE friend_id = $1
	`

	SelectUserFriendRelation = `SELECT id, user_id, friend_id, status, created_at, updated_at, deleted_at
		FROM relations
		WHERE user_id = $1 AND friend_id = $2
	`

	SelectAllRelations = `SELECT id, user_id, friend_id, status, created_at, updated_at, deleted_at FROM relations`
)
