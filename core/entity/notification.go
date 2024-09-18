package entity

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	ID               int                `json:"id"`
	PersonalEmail    string             `json:"personal"`
	SupervisorEmail  string             `json:"supervisor"`
	GoalsID          int                `json:"goalsId"`
	Message          string             `json:"message"`
	Status           string             `gorm:"type:enum('unread', 'read')" json:"status"`
	TypeNotification []TypeNotification `json:"typeNotification"`
}
type TypeNotification struct {
	gorm.Model
	ID             int    `json:"id"`
	Name           string `json:"name"`
	NotificationID int    `json:"notificationId"`
}
type NotificationRespon struct {
	ID               int                      `json:"id"`
	PersonalEmail    string                   `json:"personal"`
	SupervisorEmail  string                   `json:"supervisor"`
	GoalsID          int                      `json:"goalsId"`
	Message          string                   `json:"message"`
	Status           string                   `gorm:"type:enum('unread', 'read')" json:"status"`
	TypeNotification []TypeNotificationRespon `json:"typeNotification"`
}

type TypeNotificationRespon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	NotificationID int    `json:"notificationId"`
}
