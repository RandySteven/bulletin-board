package db

import (
	"context"
	"task_mission/entities/models"
)

func (r *Repositories) SeedRole(ctx context.Context) {
	userRole := make(map[string]uint64)
	userRole["admin"] = 1
	userRole["user"] = 2
	userRole["super-user"] = 3
	for key, value := range userRole {
		role := &models.Role{ID: value, Role: key}
		r.RoleRepository.Save(ctx, role)
	}
}

func (r *Repositories) SeedCategory(ctx context.Context) {
	categories := []models.Category{
		{
			Category: `Money`,
		},
		{
			Category: `Card`,
		},
		{
			Category: `Coupon`,
		},
	}

	for _, category := range categories {
		r.CategoryRepository.Save(ctx, &category)
	}
}