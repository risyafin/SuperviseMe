package module

import (
	"errors"
	"superviseMe/core/entity"
	"superviseMe/core/repository"
	"time"
)

type GoalsUseCase interface {
	CreateGoals(goals *entity.Goals) (*entity.Goals, error) 
	// GetGoalsByUserID(userID string) (*entity.Goals, error)
	GetGoalsByGmail(personalGmail string) (*entity.Goals, error)
	// GetGoals() (*entity.Goals, error)
	// UpdateGoals(id string, goals *entity.Goals) error
	// DeleteGoals(id string, isAvtive string) (*entity.Goals, error)
}

type goalsUseCase struct {
	goalsRepository repository.GoalsRepository
}

func NewGoalsUseCase(goalsRepository repository.GoalsRepository) GoalsUseCase {
	return &goalsUseCase{goalsRepository: goalsRepository}
}

//	func (e *goalsUseCase) GetGoals() (*entity.Goals, error) {
//		return e.goalsRepository.GetGoals()
//	}
//
//	func (e *goalsUseCase) GetGoalsByUserID(userID string) (*entity.Goals, error) {
//		return e.goalsRepository.GetGoalsByUserID(userID)
//	}
func (e *goalsUseCase) GetGoalsByGmail(personalGmail string) (*entity.Goals, error) {
	return e.goalsRepository.GetGoalsByGmail(personalGmail)
}
func (e *goalsUseCase) CreateGoals(goals *entity.Goals) (*entity.Goals, error)  {
	if goals.IsActive == "" {
		goals.IsActive = "1"
	}
	if goals.GoalStatus == "" {
		goals.GoalStatus = "progres"
	}
	if goals.PersonalGmail == *goals.SupervisorGmail {
		return nil, errors.New("tidak bisa menjadi supervisor di goal sendiri")
	}
	goals.Status = "requested"
	goals.NilaiProgres = 0
	goals.RequestedAt = time.Now()
	return e.goalsRepository.CreateGoals(goals)
}

// func (e *goalsUseCase) UpdateGoals(id string, goals *entity.Goals) error {
// 	return e.goalsRepository.UpdateGoals(id, goals)
// }
// func (e *goalsUseCase) DeleteGoals(id string, isActive string) (*entity.Goals, error) {
// 	goals, err := e.GetGoalsByID(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if isActive == "0" {
// 		if goals.IsActive == "0" {
// 			return nil, errors.New("article has been deleted")
// 		} else if goals.IsActive != "0" {
// 			goals.IsActive = "0"
// 		}
// 	} else {
// 		return nil, errors.New("you must enter the keyword '0' to delete the goals")
// 	}
// 	if err := e.goalsRepository.DeleteGoals(id, goals); err != nil {
// 		return nil, errors.New("goals cannot deleted")
// 	}
// 	return goals, nil
// }
