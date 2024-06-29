package postgres_repositories

import (
	"context"
	"database/sql"
	"task_mission/entities/models"
	"task_mission/interfaces/repositories"
	"task_mission/queries"
	"task_mission/utils"
)

type (
	chatRepository struct {
		db *sql.DB
	}

	roomRepository struct {
		db *sql.DB
	}
)

func (r *roomRepository) CheckExistsRoom(ctx context.Context, room models.Room) (exist bool, err error) {
	err = r.db.QueryRowContext(ctx, queries.SelectExistRoomForUser.ToString(), room.UserID1, room.UserID2).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (r *roomRepository) Save(ctx context.Context, request *models.Room) (result *uint64, err error) {
	return utils.Save[models.Room](ctx, r.db, queries.InsertRoom, request)
}

func (r *roomRepository) FindAll(ctx context.Context) (result []*models.Room, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *roomRepository) Find(ctx context.Context, id uint64) (result *models.Room, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *roomRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r *roomRepository) Update(ctx context.Context, request *models.Room) (result *models.Room, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *chatRepository) Save(ctx context.Context, request *models.Chat) (result *uint64, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *chatRepository) FindAll(ctx context.Context) (result []*models.Chat, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *chatRepository) Find(ctx context.Context, id uint64) (result *models.Chat, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *chatRepository) Delete(ctx context.Context, id uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (c *chatRepository) Update(ctx context.Context, request *models.Chat) (result *models.Chat, err error) {
	//TODO implement me
	panic("implement me")
}

func NewChatRepository(db *sql.DB) *chatRepository {
	return &chatRepository{
		db: db,
	}
}

func NewRoomRepository(db *sql.DB) *roomRepository {
	return &roomRepository{
		db: db,
	}
}

var (
	_ repositories.IChatRepository = &chatRepository{}
	_ repositories.IRoomRepository = &roomRepository{}
)
