package repository

import "superviseMe/core/entity"

type ListRepository interface {
	GetList(gmailPersonal string) (*entity.List, error)
}
