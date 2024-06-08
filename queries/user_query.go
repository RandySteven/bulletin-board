package queries

const (
	InsertUser                GoQuery = `INSERT INTO users (name, user_name, date_of_birth, gender) VALUES ($1, $2, $3, $4) RETURNING id`
	SelectUserByID                    = `SELECT id, name, user_name, date_of_birth, gender, is_verified, created_at, updated_at, deleted_at FROM users WHERE id = $1`
	SelectAllUser                     = `SELECT id, name, user_name, date_of_birth, gender, is_verified, created_at, updated_at, deleted_at FROM users`
	DeleteUser                        = `UPDATE users SET deleted_at = NOW() WHERE id = $1`
	UpdateUser                        = `UPDATE users SET name=$1, user_name=$2, date_of_birth=$3, gender=$4, is_verified=$5, updated_at=NOW() WHERE id = $6`
	SelectUserVerifyCondition         = `SELECT is_verified FROM users WHERE id = $1 AND is_verified=$2`
)
