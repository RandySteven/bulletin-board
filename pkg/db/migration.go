package db

import (
	"context"
	"log"
	"task_mission/queries"
)

func initTableMigration() []queries.TableMigration {
	return []queries.TableMigration{
		queries.UserTableMigration,
		queries.UserProfileMigration,
		queries.RoleMigration,
		queries.UserRolesMigration,
		queries.TaskMigration,
		queries.CategoryMigration,
		queries.RewardMigration,
		queries.RewardCategoriesMigration,
		queries.TaskRewardMigration,
		queries.UserCreditsMigration,
		queries.CreditsMigration,
		queries.RelationMigration,
		queries.UserTaskMigration,
		queries.RoomMigration,
		queries.ChatMigration,
	}
}

func (repo *Repositories) Migration(ctx context.Context) {
	migrations := initTableMigration()

	for _, migration := range migrations {
		_, err := repo.db.ExecContext(ctx, migration.ToString())
		if err != nil {
			log.Fatalf("Error in migration : %s \n", err.Error())
			return
		}
	}
}
