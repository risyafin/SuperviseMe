package module

import (
	"superviseMe/core/entity"
	"superviseMe/core/repository"
)

type NotificationUsecase interface {
	// GetNotification(personal string, supervisor string) ([]entity.Notification, error)
	UpdateNotificationStatusAndFetch(email string) ([]entity.NotificationRespon, error)
	// UpdateNotification(message, status, email string) error
}

type notificationUsecase struct {
	notificationRepository repository.NotificationRepository
}

func NewNotificationUseCase(notificationrepository repository.NotificationRepository) NotificationUsecase {
	return &notificationUsecase{notificationRepository: notificationrepository}
}

// func (e *notificationUsecase) UpdateNotification(message, status, email string) error {
// 	return e.notificationRepository.UpdateNotification(message,status,email)
// }

func (u *notificationUsecase) UpdateNotificationStatusAndFetch(email string) ([]entity.NotificationRespon, error) {
	notifications, err := u.notificationRepository.UpdateStatusAndFetchAll(email)
	if err != nil {
		return nil, err
	}

	// Convert entity.User to entity.UserResponse
	var typeNotificationRespon []entity.TypeNotification
	for _, typeNotification := range typeNotificationRespon {
		typeNotificationRespon = append(typeNotificationRespon, entity.TypeNotification{
			ID:             typeNotification.ID,
			Name:           typeNotification.Name,
			NotificationID: typeNotification.ID,
		})
	}

	var notificationResponses []entity.NotificationRespon
	for _, notification := range notifications {

		notificationResponses = append(notificationResponses, entity.NotificationRespon{
			ID:              notification.ID,
			PersonalEmail:   notification.PersonalEmail,
			SupervisorEmail: notification.SupervisorEmail,
			GoalsID:         notification.GoalsID,
			Message:         notification.Message,
			Status:          notification.Status,
		})
	}

	return notificationResponses, nil
}

// func (e *notificationUsecase) GetNotification(personal string, supervisor string) ([]entity.Notification, error) {
// 	status := "read"
// 	err := e.notificationRepository.UpdateNotificationStatus(personal, supervisor, status)
// 	if err != nil {
// 		return nil, err
// 	}
// 	notification, err := e.notificationRepository.GetAllNotification(personal, supervisor)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return notification, err
// }
