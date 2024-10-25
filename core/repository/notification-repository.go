package repository

import "superviseMe/core/entity"

type NotificationRepository interface {
	CreateNotification(notification *entity.Notification) error
	UpdateNotification(message, status, email string) error
	UpdateStatusAndFetchAll(email string) ([]entity.Notification, error)
}
