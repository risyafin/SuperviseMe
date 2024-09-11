package userrepository

import (
	"superviseMe/core/entity"
	userRepo "superviseMe/core/repository"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) userRepo.UserResitory {
	return &repository{
		DB: db,
	}
}

func (r *repository) Save(user *entity.User) error {
	return r.DB.Save(user).Error
}

func (r *repository) UpdateName(name string, email string) error {
	return r.DB.Model(&entity.User{}).Where("email = ?", email).Update("name", name).Error
}

func (r *repository) FindByID(id string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r repository) GetGoalsBySuperviseeUser(gmail string) (*entity.User, error) {
	var (
		goals *entity.User
		db    = r.DB
	)

	db = db.Preload("Comment")
	db = db.Preload("PersonalGoals", func(db *gorm.DB) *gorm.DB {
		return db.Limit(2)
	})
	db = db.Preload("SupervisorGoals", func(db *gorm.DB) *gorm.DB {
		return db.Limit(2)
	})
	db = db.Preload("ActivityLog")

	err := db.Where("email =?", gmail).First(&goals).Error
	return goals, err
}

func (r repository) GetGoalSupervisor(email string) (*entity.User, error) {
	var (
		goals *entity.User
		db    = r.DB
	)
	db = db.Preload("SupervisorGoals", func(db *gorm.DB) *gorm.DB {
		return db.Limit(3)
	})

	err := db.Where("email =?", email).First(&goals).Error
	return goals, err
}

func (r repository) GetGoalPersonal(email string) (*entity.User, error) {
	var (
		goals *entity.User
		db    = r.DB
	)
	db = db.Preload("PersonalGoals", func(db *gorm.DB) *gorm.DB {
		return db.Limit(3)
	})

	err := db.Where("email =?", email).First(&goals).Error
	return goals, err
}

func (r repository) GetUserByGmail(email string) (*entity.User, error) {
	var user *entity.User
	result := r.DB.Where("email =? ", email).First(&user)
	return user, result.Error
}

func (r repository) Registration(user *entity.User) error {
	result := r.DB.Select("Name", "Email", "Password").Create(&user)
	return result.Error
}

func (r repository) Login(email string, password string) (*entity.User, error) {
	var user entity.User
	result := r.DB.Model(&user).Where("email =? AND password =?  ", email, password).First(&user)
	return &user, result.Error
}
