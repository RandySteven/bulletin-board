package relations

import "task_mission/entities/models"

type UserUserProfile struct {
	User    *models.User
	Profile *models.UserProfile
}
