package goalsrepository

import (
	"superviseMe/core/entity"
	goalsRepo "superviseMe/core/repository"
	"time"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewGoalsRepository(db *gorm.DB) goalsRepo.GoalsRepository {
	return &repository{
		DB: db,
	}
}

func (r repository) AcceptedSupervisor(id int, status string, accepted time.Time) error {
	updates := map[string]interface{}{
		"status":      status,
		"accepted_at": accepted,
	}

	// Update 2 field dari entitas Goals berdasarkan kondisi
	return r.DB.Model(&entity.Goals{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (r repository) RejectedSupervisor(id int, status string, reject time.Time) error {
	updates := map[string]interface{}{
		"status":      status,
		"rejected_at": reject,
	}

	// Update 2 field dari entitas Goals berdasarkan kondisi
	return r.DB.Model(&entity.Goals{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (r *repository) UpdateName(name string, email string) error {
	return r.DB.Model(&entity.User{}).Where("email = ?", email).Update("name", name).Error
}

func (r repository) GetGoalsById(id int) (*entity.Goals, error) {
	var (
		goals *entity.Goals
		db    = r.DB
	)

	err := db.Where("id = ?", id).First(&goals).Error
	return goals, err
}

// func (r repository) GetGoalsByUserID(userID string) (*entity.Goals, error) {
// 	var (
// 		goals *entity.Goals
// 		db    = r.DB
// 	)

// 	db = db.Preload("User")
// 	err := db.Where("userID = ?", userID).First(&goals).Error
// 	return goals, err
// }

// func (r repository) GetGoals() (*entity.Goals, error) {
// 	var (
// 		goals *entity.Goals
// 		db    = r.DB
// 	)

// 	db = db.Preload("User")
// 	err := db.Find(&goals).Error
// 	return goals, err
// }

func (r repository) CreateGoals(goals *entity.Goals) (*entity.Goals, error) {

	err := r.DB.Select(
		"GoalName",
		"Description",
		"PersonalGmail",
		"SupervisorGmail",
		"BackgroundColor",
		"Status",
		"NilaiProgres",
		"GoalStatus",
		"IsActive",
		"RequestedAt").Create(&goals).Error
	return goals, err
}

// func (r repository) UpdateGoals(id string, goals *entity.Goals) error {
// 	err := r.DB.Model(&entity.Goals{}).Where(id).Updates(&goals).Error
// 	return err
// }

// func (r repository) DeleteGoals(id string, goals *entity.Goals) error {
// 	err := r.DB.Save(&goals).Error
// 	return err
// }
