package responses

import "time"

type (
	RoomResponse struct {
		ID        uint64               `json:"id"`
		User1     UserRoomChatResponse `json:"user1"`
		User2     UserRoomChatResponse `json:"user2"`
		CreatedAt time.Time            `json:"created_at"`
	}

	RoomListResponse struct {
		ID   uint64               `json:"id"`
		User UserRoomChatResponse `json:"user"`
	}

	ChatResponse struct {
		Message string `json:"message"`
	}
)
