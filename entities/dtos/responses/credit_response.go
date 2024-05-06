package responses

type (
	UserCreditResponse struct {
		UserID uint64  `json:"user_id"`
		Credit float32 `json:"credit"`
	}
)
