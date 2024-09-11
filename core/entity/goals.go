package entity

import (
	"time"

	"gorm.io/gorm"
)

type Goals struct {
	gorm.Model
	ID              int     `json:"id"`
	GoalName        string  `json:"goalName"`
	Description     string  `json:"description"`
	BackgroundColor string  `json:"backgroundColor"`
	PersonalGmail   string  `gorm:"type:varchar(100);index" json:"personal"`
	SupervisorGmail string  `gorm:"type:varchar(100);index" json:"supervisor"`
	Status          string  `gorm:"type:enum('requested', 'accepted', 'rejected' )" json:"status"`
	NilaiProgres    float64 `json:"nilaiProgres"`
	GoalStatus      string  `gorm:"type:enum('progres', 'completed')" json:"goalStatus"`
	// AcceptedAt       time.Time      `json:"acceptAt"`
	// RejectedAt       time.Time      `json:"rejectedAt"`
	Requested              time.Time      `json:"requestAt"`
	IsActive               string         `gorm:"type:enum('1', '0')" json:"isActive"`
	List                   []List         `gorm:"foreignKey:GoalID;references:ID" json:"List"`
	PersonalGoalCard       []Card         `gorm:"foreignKey:GmailPersonal;references:PersonalGmail" json:"PersonalGoalCard"`
	NotificationSupervisor []Notification `gorm:"foreignKey:Email;references:SupervisorGmail" json:"notificationPersonal"`
	Notification           []Notification `json:"notification"`
}

// type RequestSupervisor struct {
// 	gorm.Model
// 	ID          int       `json:"id"`
// 	Personal    string    `gorm:"foreignKey:Personal;references:PersonalGmail" json:"personal"`
// 	Supervisor  string    `gorm:"foreignKey:Supervisor;references:PersonalGmail" json:"supervisor"`
// 	RequestedAt time.Time `json:"requestedAt"`
// }

type CreateResponGoal struct {
	ID              int    `json:"id"`
	GoalName        string `json:"goalName"`
	Supervisor      string `json:"supervisor"`
	BackgroundColor string `json:"backgroundColor"`
}

type GoalSupervisorRespons struct {
	GoalStatus      string    `json:"GoalStatus"`
	NilaiProgres    float64   `json:"nilaiProgres"`
	SupervisorGmail string    `json:"supervisor"`
	GoalName        string    `json:"goalName"`
	CreatedAt       time.Time `json:"createdAt"`
}

type GoalPersonalRespons struct {
	GoalStatus    string    `json:"GoalStatus"`
	NilaiProgres  float64   `json:"nilaiProgres"`
	PersonalGmail string    `json:"personal"`
	GoalName      string    `json:"goalName"`
	CreatedAt     time.Time `json:"createdAt"`
}

type GoalsResponseHome struct {
	ID              int       `json:"id"`
	GoalName        string    `json:"goalName"`
	Description     string    `json:"description"`
	NilaiProgres    float64   `json:"nilaiProgres"`
	GoalStatus      string    `json:"goalStatus"`
	PersonalGmail   string    `json:"personal"`
	SupervisorGmail string    `json:"supervisor"`
	CreatedAt       time.Time `json:"createdAt"`
}
