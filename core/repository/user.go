package repository

import "superviseMe/core/entity"

type UserResitory interface {
	Save(user *entity.User) error
	FindByID(id string) (*entity.User, error)
	GetGoalsBySuperviseeUser(gmail string) (*entity.User, error)
	GetGoalSupervisor(gmail string) (*entity.User, error)
	GetGoalPersonal(gmail string) (*entity.User, error)
	Login(gmail string, password string) (*entity.User, error)
	Registration(user *entity.User) error
	GetUserByGmail(gmail string) (*entity.User, error)
}
