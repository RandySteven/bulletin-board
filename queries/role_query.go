package queries

const (
	InsertRole GoQuery = `
		INSERT INTO roles (id, role)
		VALUES 
		    ($1, $2)
		RETURNING id
	`

	SelectRoleByID = `
		SELECT id, role, created_at, updated_at, deleted_at
		FROM roles
		WHERE id = $1
	`

	SelectAllRoles = `
		SELECT id, role, created_at, updated_at, deleted_at
		FROM roles
	`
)
