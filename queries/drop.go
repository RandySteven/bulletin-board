package queries

const (
	DropUserTask         DropTable = `DROP TABLE IF EXISTS user_tasks`
	DropRelation                   = `DROP TABLE IF EXISTS relations`
	DropCredits                    = `DROP TABLE IF EXISTS credits`
	DropUserCredits                = `DROP TABLE IF EXISTS user_credits`
	DropRewardCategories           = `DROP TABLE IF EXISTS reward_categories`
	DropCategories                 = `DROP TABLE IF EXISTS categories`
	DropTaskRewards                = `DROP TABLE IF EXISTS task_rewards`
	DropRewards                    = `DROP TABLE IF EXISTS rewards`
	DropTasks                      = `DROP TABLE IF EXISTS tasks`
	DropUserProfiles               = `DROP TABLE IF EXISTS user_profiles`
	DropUserRoles                  = `DROP TABLE IF EXISTS user_roles`
	DropRoles                      = `DROP TABLE IF EXISTS roles`
	DropUsers                      = `DROP TABLE IF EXISTS users`
)
