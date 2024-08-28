package module

import (
	"superviseMe/core/entity"
	"superviseMe/core/repository"
)

type ListUsecase interface {
	GetList(gmailPersonal string) (*entity.List, error)
}

type listUsecase struct {
	listRepository  repository.ListRepository
}

func NewListUseCase(listrepository repository.ListRepository) ListUsecase {
	return &listUsecase{listRepository: listrepository}
}

func (e *listUsecase) GetList(gmailPersonal string) (*entity.List, error) {
	return e.listRepository.GetList(gmailPersonal)
}
