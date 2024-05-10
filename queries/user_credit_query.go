package queries

const (
	InsertUserCredit     GoQuery = `INSERT INTO user_credits (user_id) VALUES ($1) RETURNING id`
	SelectUserCreditByID         = ``
)
