package queries

const (
	InsertChat      GoQuery = `INSERT INTO chats (room_id, user_id, message) VALUES ($1, $2, $3) RETURNING id`
	SelectAllChats  GoQuery = `SELECT * FROM chats`
	SelectChatsUser GoQuery = `SELECT * FROM chats WHERE user_id = $1`
	SelectChatsRoom GoQuery = `SELECT * FROM chats WHERE room_id = $1`
	SelectChatByID  GoQuery = `SELECT * FROM chats WHERE id = $1`
)

const (
	InsertRoom             GoQuery = `INSERT INTO rooms (user_id_1, user_id_2) VALUES ($1, $2)`
	SelectAllRooms         GoQuery = `SELECT * FROM rooms`
	SelectRoom             GoQuery = `SELECT * FROM rooms WHERE id = $1`
	SelectExistRoomForUser GoQuery = `
		SELECT EXISTS (
			SELECT *
			FROM rooms 
			WHERE 
				(user_id_1 = $1 AND user_id_2 = $2) 
			   OR 
				(user_id_1 = $2 AND user_id_2 = $1)
		)
	`
	SelectLoginUserRooms GoQuery = `
		SELECT id, user_id_1, user_id_2, created_at FROM rooms
		WHERE
		    user_id_1 = $1 OR user_id_2 = $1
	`
)
