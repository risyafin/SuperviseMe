package module

import (
	"superviseMe/core/entity"
	"superviseMe/core/repository"
)

type NotificationUsecase interface {
	GetNotification(personal string, supervisor string) (*entity.Notification, error)
}

type notificationUsecase struct {
	notificationRepository repository.NotificationRepository
}

func NewNotificationUseCase(notificationrepository repository.NotificationRepository) NotificationUsecase {
	return &notificationUsecase{notificationRepository: notificationrepository}
}

func (e *notificationUsecase) GetNotification(personal string, supervisor string) (*entity.Notification, error) {
	return e.notificationRepository.GetNotification(personal, supervisor)
}
