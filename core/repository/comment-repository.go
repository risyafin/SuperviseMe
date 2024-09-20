package repository

import "superviseMe/core/entity"

type CommentRepository interface {
	CreateComment(comment *entity.Comment) error
	// GetNotification(personal string, supervisor string) (*entity.Notification, error)
}
