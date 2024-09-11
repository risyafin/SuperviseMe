package repository

import "superviseMe/core/entity"

type NotificationRepository interface {
	CreateNotification(notification *entity.Notification) error
}
