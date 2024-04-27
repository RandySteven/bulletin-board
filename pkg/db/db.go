package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"log"
	"task_mission/interfaces/repositories"
	"task_mission/pkg/config"
	repositories2 "task_mission/repositories"
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
	UnitOfWork               repositories.UnitOfWork
	db                       *sql.DB
}

func NewRepositories(config *config.Config) (*Repositories, error) {
	conn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=require",
		config.Postgres.DbUser,
		config.Postgres.DbPass,
		config.Postgres.Host,
		config.Postgres.DbName,
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
		UserRepository:           repositories2.NewUserRepository(db),
		UserProfileRepository:    repositories2.NewUserProfileRepository(db),
		RoleRepository:           repositories2.NewRoleRepository(db),
		CategoryRepository:       repositories2.NewCategoryRepository(db),
		TaskRepository:           repositories2.NewTaskRepository(db),
		RewardRepository:         repositories2.NewRewardRepository(db),
		RewardCategoryRepository: repositories2.NewRewardCategoryRepository(db),
		TaskRewardRepository:     repositories2.NewTaskRewardRepository(db),
		UserCreditRepository:     repositories2.NewUserCreditRepository(db),
		UserRoleRepository:       repositories2.NewUserRoleRepository(db),
		UserTaskRepository:       repositories2.NewUserTaskRepository(db),
		RelationRepository:       repositories2.NewRelationRepository(db),
		UnitOfWork:               repositories2.NewUnitOfWork(db),
		db:                       db,
	}, nil
}
