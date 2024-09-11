package notificationrepository

import (
	"fmt"
	"superviseMe/core/entity"
	notifRepo "superviseMe/core/repository"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) notifRepo.NotificationRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateNotification(notification *entity.Notification) error {
	err := r.DB.Create(&notification).Error
	if err != nil {
		// Log the error (opsional)
		fmt.Println("Failed to create notification:", err)
		return err
	}
	return nil
}
