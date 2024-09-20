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

func (r *repository) GetNotification(personal string, supervisor string) (*entity.Notification, error) {
	var (
		notification *entity.Notification
		db           = r.DB
	)

	err := db.Where("personal_email = ? OR supervisor_email = ?", personal, supervisor).First(&notification).Error
	return notification, err
}
