package repositories

import "task_mission/entities/dtos/relations"

type IUserUserProfileRepository interface {
	IRepository[relations.UserUserProfile]
}
