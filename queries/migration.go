package queries

const (
	UserTableMigration TableMigration = `
		CREATE TABLE IF NOT EXISTS users (
		    id BIGSERIAL PRIMARY KEY,
		    name VARCHAR NOT NULL,
		    user_name VARCHAR NOT NULL UNIQUE,
		    date_of_birth DATE NOT NULL,
		    gender VARCHAR NOT NULL,
		    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
			deleted_at TIMESTAMP NULL		    
		)
	`

	RoleMigration = `
		CREATE TABLE IF NOT EXISTS roles (
		    id BIGSERIAL PRIMARY KEY,
		    role VARCHAR NOT NULL UNIQUE,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
			deleted_at TIMESTAMP NULL
		)
	`

	UserRolesMigration = `
		CREATE TABLE IF NOT EXISTS user_roles (
			id BIGSERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			role_id BIGINT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP NULL,
		    FOREIGN KEY (user_id) REFERENCES users(id),
		    FOREIGN KEY (role_id) REFERENCES roles(id)
		)
	`

	UserProfileMigration = `
		CREATE TABLE IF NOT EXISTS user_profiles(
		    id BIGSERIAL PRIMARY KEY,
		    email VARCHAR UNIQUE NOT NULL,
		    password VARCHAR NOT NULL,
		    image VARCHAR NOT NULL, 
		    user_id BIGINT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP NULL,
		    FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`

	TaskMigration = `
		CREATE TABLE IF NOT EXISTS tasks (
		    id BIGSERIAL PRIMARY KEY,
		    title VARCHAR NOT NULL,
		    description VARCHAR NOT NULL,
		    image VARCHAR NOT NULL,
		    expired_date DATE NOT NULL,
		    status VARCHAR NOT NULL,
		    user_id BIGINT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP NULL,
		    FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`

	RewardMigration = `
		CREATE TABLE IF NOT EXISTS rewards (
		    id BIGSERIAL PRIMARY KEY,
		    name VARCHAR NOT NULL,
		    description VARCHAR NOT NULL,
		    image VARCHAR NOT NULL,
		    user_id BIGINT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP NULL,
		    FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`

	TaskRewardMigration = `
		CREATE TABLE IF NOT EXISTS task_rewards (
			id BIGSERIAL PRIMARY KEY,
			task_id BIGINT NOT NULL,
			reward_id BIGINT UNIQUE,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP NULL,
		    FOREIGN KEY (task_id) REFERENCES tasks(id),
		    FOREIGN KEY (reward_id) REFERENCES rewards(id)
		)
	`

	CategoryMigration = `
		CREATE TABLE IF NOT EXISTS categories (
		    id BIGSERIAL PRIMARY KEY,
		    category VARCHAR NOT NULL UNIQUE,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP NULL
		)
	`

	RewardCategoriesMigration = `
		CREATE TABLE IF NOT EXISTS reward_categories (
		    id BIGSERIAL PRIMARY KEY,
		    reward_id BIGINT NOT NULL,
		    category_id BIGINT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP NULL,
		    FOREIGN KEY (reward_id) REFERENCES rewards (id),
		    FOREIGN KEY (category_id) REFERENCES categories (id)
		)
	`

	UserCreditsMigration = `
		CREATE TABLE IF NOT EXISTS user_credits (
		    id BIGSERIAL PRIMARY KEY,
		    user_id BIGINT NOT NULL,
		    credit FLOAT NOT NULL DEFAULT 0,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP NULL,
		    FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`

	CreditsMigration = `
		CREATE TABLE IF NOT EXISTS credits (
		    id BIGSERIAL PRIMARY KEY,
		    to_id BIGINT NOT NULL,
		    from_id BIGINT NOT NULL,
		    credit FLOAT NOT NULL,
		    description VARCHAR NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP NULL,
		    FOREIGN KEY (to_id) REFERENCES users(id),
		    FOREIGN KEY (from_id) REFERENCES users(id)
		)
	`

	RelationMigration = `
		CREATE TABLE IF NOT EXISTS relations (
		    id BIGSERIAL PRIMARY KEY,
		    user_id BIGINT NOT NULL,
		    friend_id BIGINT NOT NULL,
		    status VARCHAR NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP NULL,
		    FOREIGN KEY(user_id) REFERENCES users(id),
		    FOREIGN KEY(friend_id) REFERENCES users(id)
		)
	`

	UserTaskMigration = `
		CREATE TABLE IF NOT EXISTS user_tasks (
			id BIGSERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			task_id BIGINT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP NULL,
		    FOREIGN KEY(user_id) REFERENCES users (id),
		    FOREIGN KEY(task_id) REFERENCES tasks (id)
		)
	`

	RoomMigration = `
		CREATE TABLE IF NOT EXISTS rooms (
		    id BIGSERIAL PRIMARY KEY,
		    user_id_1 BIGINT NOT NULL,
		    user_id_2 BIGINT NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP NULL,
		    FOREIGN KEY(user_id_1) REFERENCES users(id),
		    FOREIGN KEY(user_id_2) REFERENCES users(id)
		)
	`

	ChatMigration = `
		CREATE TABLE IF NOT EXISTS chats (
			id BIGSERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			room_id BIGINT NOT NULL,
			message VARCHAR NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP NULL,
		    FOREIGN KEY(user_id) REFERENCES users(id),
		    FOREIGN KEY(room_id) REFERENCES rooms(id)
		)
	`
)
