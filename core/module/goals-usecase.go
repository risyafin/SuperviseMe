package module

import (
	"errors"
	"fmt"
	"superviseMe/core/entity"
	"superviseMe/core/repository"
	"time"
)

type GoalsUseCase interface {
	CreateGoals(goals *entity.Goals) (*entity.Goals, error)
	GetGoalsByGmail(personalGmail string) (*entity.Goals, error)
	RequestSupervisor(personalEmail string, supervisorEmail string) error
	AcceptSupervisorRequest(gmail string, supervisorEmail string) error
	RejectSupervisorRequest(gmail string, supervisorEmail string) error
	// GetGoalsByUserID(userID string) (*entity.Goals, error)
	// GetGoals() (*entity.Goals, error)
	// UpdateGoals(id string, goals *entity.Goals) error
	// DeleteGoals(id string, isAvtive string) (*entity.Goals, error)
}

type goalsUseCase struct {
	goalsRepository  repository.GoalsRepository
	notificationRepo repository.NotificationRepository
}

func NewGoalsUseCase(goalsRepository repository.GoalsRepository, notificationRepo repository.NotificationRepository) GoalsUseCase {
	return &goalsUseCase{
		goalsRepository:  goalsRepository,
		notificationRepo: notificationRepo,
	}
}

//	func (e *goalsUseCase) GetGoals() (*entity.Goals, error) {
//		return e.goalsRepository.GetGoals()
//	}
//
//	func (e *goalsUseCase) GetGoalsByUserID(userID string) (*entity.Goals, error) {
//		return e.goalsRepository.GetGoalsByUserID(userID)
//	}

func (uc *goalsUseCase) RequestSupervisor(personalEmail string, supervisorEmail string) error {
	goal, err := uc.goalsRepository.GetGoalsByGmail(personalEmail)
	if err != nil {
		return err
	}

	goal.SupervisorGmail = supervisorEmail
	goal.Status = "requested"
	// goal.RequestedAt = time.Now()

	err = uc.goalsRepository.RequestSupervisor(goal)
	if err != nil {
		return err
	}

	// Create a notification for the supervisor
	message := fmt.Sprintf("You have a new request to supervise the goal: %s", goal.GoalName)
	notification := &entity.Notification{
		Email:   goal.PersonalGmail, // Assuming `UserId` refers to the supervisor's user ID
		GoalsID: goal.ID,
		Message: message,
		Status:  "unread",
	}

	err = uc.notificationRepo.CreateNotification(notification)
	if err != nil {
		return err
	}

	return nil
}

func (uc *goalsUseCase) AcceptSupervisorRequest(gmail string, supervisorEmail string) error {
	goal, err := uc.goalsRepository.GetGoalsByGmail(gmail)
	if err != nil {
		return err
	}

	if goal.SupervisorGmail != supervisorEmail || goal.Status != "requested" {
		return errors.New("invalid supervisor email or request not pending")
	}

	goal.Status = "accepted"
	// goal.AcceptedAt = time.Now()

	return uc.goalsRepository.RequestSupervisor(goal)
}

func (uc *goalsUseCase) RejectSupervisorRequest(gmail string, supervisorEmail string) error {
	goal, err := uc.goalsRepository.GetGoalsByGmail(gmail)
	if err != nil {
		return err
	}

	if goal.SupervisorGmail != supervisorEmail || goal.Status != "requested" {
		return errors.New("invalid supervisor email or request not pending")
	}

	goal.Status = "rejected"
	// goal.RejectedAt = time.Now()

	return uc.goalsRepository.RequestSupervisor(goal)
}

func (e *goalsUseCase) GetGoalsByGmail(personalGmail string) (*entity.Goals, error) {
	return e.goalsRepository.GetGoalsByGmail(personalGmail)
}

func (e *goalsUseCase) CreateGoals(goals *entity.Goals) (*entity.Goals, error) {
	if e.goalsRepository == nil {
		return nil, errors.New("goalsRepository is not initialized")
	}
	// if e.notificationRepo == nil {
	//     return nil, errors.New("notificationRepo is not initialized")
	// }

	if goals.IsActive == "" {
		goals.IsActive = "1"
	}
	if goals.GoalStatus == "" {
		goals.GoalStatus = "progres"
	}
	if goals.PersonalGmail == goals.SupervisorGmail {
		return nil, errors.New("tidak bisa menjadi supervisor di goal sendiri")
	}

	goals.Status = "requested"
	goals.NilaiProgres = 0
	goals.Requested = time.Now()

	// err := e.goalsRepository.RequestSupervisor(goals)
	// if err != nil {
	// 	return nil, err
	// }
	fmt.Println("ini usecase", goals.ID)
	goal, err := e.goalsRepository.CreateGoals(goals)
	if err != nil {
		return goal, err
	}

	// Create a notification for the supervisor
	message := fmt.Sprintf("You have a new request to supervise the goal: %s", goal.GoalName)
	notification := &entity.Notification{
		Email:   goal.SupervisorGmail, // Assuming `UserId` refers to the supervisor's user ID
		GoalsID: goal.ID,
		Message: message,
		Status:  "unread",
	}

	fmt.Println("ini dibawah", goal.ID)
	errn := e.notificationRepo.CreateNotification(notification)
	if errn != nil {
		fmt.Println("ini error", errn)
	}

	return goal, err
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
