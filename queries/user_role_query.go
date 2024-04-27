package queries

const (
	InsertUserRole GoQuery = `
		INSERT INTO user_roles (user_id, role_id) VALUES 
		                                              ($1, $2)
		RETURNING id
	`

	SelectUserRoleByUserID = `
		SELECT id, user_id, role_id, created_at, updated_at, deleted_at 
		FROM user_roles
		WHERE user_id = $1
	`
)
