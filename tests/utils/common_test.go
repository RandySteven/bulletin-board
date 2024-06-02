package utils_test

import (
	"github.com/stretchr/testify/assert"
	"task_mission/entities/models"
	"task_mission/utils"
	"testing"
	"time"
)

func TestUserFullName(t *testing.T) {
	firstName := "Randy"
	lastName := "Steven"
	expectedResult := firstName + " " + lastName

	actualResutl := utils.UserFullName(firstName, lastName)

	assert.Equal(t, expectedResult, actualResutl)
}

func TestCreditAverage(t *testing.T) {
	credits := []*models.Credit{
		&models.Credit{
			ID:          1,
			FromID:      1,
			ToID:        2,
			Credit:      10.0,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		&models.Credit{
			ID:          2,
			FromID:      1,
			ToID:        2,
			Credit:      10.0,
			Description: "",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	creditLen := len(credits)
	creditSum := float32(0)
	for _, credit := range credits {
		creditSum += credit.Credit
	}
	expectedResult := creditSum / float32(creditLen)

	actualResult := utils.CreditsAverage(credits)

	assert.Equal(t, expectedResult, actualResult)
}
