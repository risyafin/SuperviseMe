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

func (r *repository) UpdateNotification(message, status, email string) error {
	return r.DB.Model(&entity.Notification{}).Where("personal_email = ? OR supervisor_email = ?", email, email).Updates(entity.Notification{Message: message, Status: status}).Error
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
func (r *repository) UpdateStatusAndFetchAll(email string) ([]entity.Notification, error) {
	err := r.DB.Model(&entity.Notification{}).Where("personal_email = ? OR supervisor_email = ?", email, email).Update("status", "read").Error
	if err != nil {
		return nil, err
	}

	// Setelah update, ambil semua data user
	var notification []entity.Notification
	err = r.DB.Where("personal_email = ? OR supervisor_email = ?", email, email).Find(&notification).Error
	if err != nil {
		return nil, err
	}

	return notification, nil
}
