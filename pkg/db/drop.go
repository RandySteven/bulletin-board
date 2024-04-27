package db

import (
	"context"
	"log"
	"task_mission/queries"
)

func initDropTables() []queries.DropTable {
	return []queries.DropTable{
		queries.DropUserTask,
		queries.DropRelation,
		queries.DropCredits,
		queries.DropUserCredits,
		queries.DropRewardCategories,
		queries.DropCategories,
		queries.DropTaskRewards,
		queries.DropRewards,
		queries.DropTasks,
		queries.DropUserProfiles,
		queries.DropUserRoles,
		queries.DropRoles,
		queries.DropUsers,
	}
}

func (r *Repositories) DropTable(ctx context.Context) {
	dropTables := initDropTables()

	for _, drop := range dropTables {
		_, err := r.db.ExecContext(ctx, drop.ToString())
		if err != nil {
			log.Fatalf("Error drop table : %s ", err.Error())
			return
		}
	}
}
