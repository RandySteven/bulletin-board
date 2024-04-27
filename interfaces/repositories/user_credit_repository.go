package repositories

import "task_mission/entities/models"

type IUserCreditRepository interface {
	IRepository[models.UserCredit]
}
