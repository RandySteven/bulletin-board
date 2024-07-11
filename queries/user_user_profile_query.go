package queries

const (
	SelectAllUserUserProfiles = `
		SELECT u.id,
		       u.name,
		       u.user_name,
		       u.date_of_birth,
		       u.gender,
		       u.is_verified,
		       up.email,
		       up.image
		FROM users u INNER JOIN user_profiles up ON u.id = up.user_id
		WHERE u.deleted_at IS NULL
	`
)
