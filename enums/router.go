package enums

type RouterPrefix string

func (r RouterPrefix) ToString() string {
	return string(r)
}

const (
	AuthRouter     RouterPrefix = `/auth`
	UserRouter                  = `/users`
	TaskRouter     RouterPrefix = `/tasks`
	RelationRouter RouterPrefix = `/relations`
	RewardRouter   RouterPrefix = `/rewards`
	CategoryRouter RouterPrefix = `/categories`
)
