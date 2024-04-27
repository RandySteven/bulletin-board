package enums

type RouterPrefix string

func (r RouterPrefix) ToString() string {
	return string(r)
}

const (
	AuthRouter     RouterPrefix = `/auth`
	UserRouter                  = `/users`
	TaskRouter                  = `/tasks`
	RelationRouter              = `/relations`
)
