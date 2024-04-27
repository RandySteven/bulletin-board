package responses

import "time"

type CreateFriendResponse struct {
	ID         uint64     `json:"id"`
	UserID     uint64     `json:"user_id"`
	UserName   string     `json:"user_name"`
	FriendID   uint64     `json:"friend_id"`
	FriendName string     `json:"friend_name"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
