package repository

import (
	"superviseMe/core/entity"
	"time"
)

type GoalsRepository interface {
	CreateGoals(goals *entity.Goals) (*entity.Goals, error)
	GetGoalsById(id int) (*entity.Goals, error)
	AcceptedSupervisor(id int, status string, accepted time.Time) error
	RejectedSupervisor(id int, status string, reject time.Time) error

	// GetGoals() (*entity.Goals, error)
	// GetGoalsByUserID(userID string) (*entity.Goals, error)
	// UpdateGoals(id string, goals *entity.Goals) error
	// DeleteGoals(id string, goals *entity.Goals) error
}
