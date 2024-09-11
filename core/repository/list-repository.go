package repository

import "superviseMe/core/entity"

type ListRepository interface {
	GetList(emailPersonal string) (*entity.List, error)
}
