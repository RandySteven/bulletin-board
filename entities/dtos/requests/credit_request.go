package requests

type CreditRequest struct {
	ToUserID    uint64  `json:"to_user_id"`
	Credit      float64 `json:"credit"`
	Description string  `json:"description"`
}
