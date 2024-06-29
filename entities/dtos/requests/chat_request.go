package requests

type (
	RoomRequest struct {
		UserID uint64 `json:"user_id"`
	}

	ChatRequest struct {
		RoomID  uint64 `json:"room_id"`
		Message string `json:"message"`
	}
)
