package module

import (
	"superviseMe/core/entity"
	"superviseMe/core/repository"
)

type CommentUsecase interface {
	CreateComment(cardID, userID int, message string) error
}

type commentUsecase struct {
	commentRepository repository.CommentRepository
}

func NewCommentUsecase(commentRepository repository.CommentRepository) CommentUsecase {
	return &commentUsecase{commentRepository: commentRepository}
}

func (u *commentUsecase) CreateComment(cardID, userID int, message string) error {

	comment := &entity.Comment{
		CardID:  cardID,
		UserID:  userID,
		Message: message,
	}
	return u.commentRepository.CreateComment(comment)
}
