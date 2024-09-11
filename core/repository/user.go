package repository

import "superviseMe/core/entity"

type UserResitory interface {
	Save(user *entity.User) error
	FindByID(id string) (*entity.User, error)
	GetGoalsBySuperviseeUser(email string) (*entity.User, error)
	GetGoalSupervisor(email string) (*entity.User, error)
	GetGoalPersonal(email string) (*entity.User, error)
	Login(email string, password string) (*entity.User, error)
	Registration(user *entity.User) error
	GetUserByGmail(email string) (*entity.User, error)
	UpdateName(name string, email string) error 
}
