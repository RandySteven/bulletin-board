package requests

type (
	RoomRequest struct {
		UserID uint64 `json:"user_id"`
	}

	ChatRequest struct {
		Message string `json:"message"`
	}
)
