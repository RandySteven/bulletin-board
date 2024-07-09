package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"log"
	"task_mission/interfaces/repositories"
	"task_mission/pkg/config"
	"task_mission/repositories/postgres"
	"time"
)

type Repositories struct {
	UserRepository           repositories.IUserRepository
	RoleRepository           repositories.IRoleRepository
	CategoryRepository       repositories.ICategoryRepository
	UserProfileRepository    repositories.IUserProfileRepository
	UserRoleRepository       repositories.IUserRoleRepository
	TaskRepository           repositories.ITaskRepository
	RewardRepository         repositories.IRewardRepository
	TaskRewardRepository     repositories.ITaskRewardRepository
	RewardCategoryRepository repositories.IRewardCategoryRepository
	UserCreditRepository     repositories.IUserCreditRepository
	UserTaskRepository       repositories.IUserTaskRepository
	RelationRepository       repositories.IRelationRepository
	CreditRepository         repositories.ICreditRepository
	RoomRepository           repositories.IRoomRepository
	ChatRepository           repositories.IChatRepository
	UnitOfWork               repositories.UnitOfWork
	db                       *sql.DB
}

func NewRepositories(config *config.Config) (*Repositories, error) {
	if config.Postgres.Port != "" {
		config.Postgres.Host = config.Postgres.Host + ":" + config.Postgres.Port
	}
	conn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s",
		config.Postgres.DbUser,
		config.Postgres.DbPass,
		config.Postgres.Host,
		config.Postgres.DbName,
		config.Postgres.SslMode,
	)
	log.Println(conn)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(8)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(8 * time.Minute)
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return &Repositories{
		UnitOfWork:               postgres_repositories.NewUnitOfWork(db),
		UserRepository:           postgres_repositories.NewUserRepository(db),
		UserProfileRepository:    postgres_repositories.NewUserProfileRepository(db),
		RoleRepository:           postgres_repositories.NewRoleRepository(db),
		CategoryRepository:       postgres_repositories.NewCategoryRepository(db),
		TaskRepository:           postgres_repositories.NewTaskRepository(db),
		RewardRepository:         postgres_repositories.NewRewardRepository(db),
		RewardCategoryRepository: postgres_repositories.NewRewardCategoryRepository(db),
		TaskRewardRepository:     postgres_repositories.NewTaskRewardRepository(db),
		UserCreditRepository:     postgres_repositories.NewUserCreditRepository(db),
		UserRoleRepository:       postgres_repositories.NewUserRoleRepository(db),
		UserTaskRepository:       postgres_repositories.NewUserTaskRepository(db),
		RelationRepository:       postgres_repositories.NewRelationRepository(db),
		CreditRepository:         postgres_repositories.NewCreditRepository(db),
		RoomRepository:           postgres_repositories.NewRoomRepository(db),
		ChatRepository:           postgres_repositories.NewChatRepository(db),
		db:                       db,
	}, nil
}
