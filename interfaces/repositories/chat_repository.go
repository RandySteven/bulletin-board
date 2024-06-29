package repositories

import (
	"context"
	"task_mission/entities/models"
)

type (
	IChatRepository interface {
		IRepository[models.Chat]
	}

	IRoomRepository interface {
		IRepository[models.Room]
		CheckExistsRoom(ctx context.Context, room *models.Room) (exist bool, err error)
		FindAllUserRooms(ctx context.Context) ([]*models.Room, error)
	}
)
