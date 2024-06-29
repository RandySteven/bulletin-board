package responses

import "time"

type (
	RoomResponse struct {
		ID        uint64                `json:"id"`
		User1     *UserRoomChatResponse `json:"user1"`
		User2     *UserRoomChatResponse `json:"user2"`
		CreatedAt time.Time             `json:"created_at"`
	}

	RoomListResponse struct {
		ID    uint64                `json:"id"`
		User1 *UserRoomChatResponse `json:"user1"`
		User2 *UserRoomChatResponse `json:"user2"`
	}

	ChatResponse struct {
		UserName  string    `json:"user_name"`
		Message   string    `json:"message"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
