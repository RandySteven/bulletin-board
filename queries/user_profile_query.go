package queries

const (
	InsertUserProfile GoQuery = `
		INSERT INTO user_profiles (email, password, image, user_id)
		VALUES 
		    ($1, $2, $3, $4)
		RETURNING id
	`
	SelectAllUserProfiles = `
		SELECT id, email, password, image, user_id, created_at, updated_at, deleted_at
		FROM user_profiles
	`
	SelectUserProfileByID = `
		SELECT id, email, password, image, user_id, created_at, updated_at, deleted_at
		FROM user_profiles
		WHERE id = $1
	`
	SelectUserProfileByEmail = `
		SELECT id, email, password, image, user_id, created_at, updated_at, deleted_at
		FROM user_profiles
		WHERE email = $1
	`
	SelectUserProfileByUserID = `
		SELECT id, email, password, image, user_id, created_at, updated_at, deleted_at
		FROM user_profiles
		WHERE user_id = $1
	`
)
