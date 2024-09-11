package repository

import "superviseMe/core/entity"

type GoalsRepository interface {
	CreateGoals(goals *entity.Goals) (*entity.Goals, error)
	GetGoalsByGmail(personalGmail string) (*entity.Goals, error)
	RequestSupervisor(goal *entity.Goals) error

	// GetGoals() (*entity.Goals, error)
	// GetGoalsByUserID(userID string) (*entity.Goals, error)
	// UpdateGoals(id string, goals *entity.Goals) error
	// DeleteGoals(id string, goals *entity.Goals) error
}
