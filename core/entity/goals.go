package entity

import (
	"time"

	"gorm.io/gorm"
)

type Goals struct {
	gorm.Model
	ID               int            `json:"id"`
	GoalName         string         `json:"goalName"`
	Description      string         `json:"description"`
	BackgroundColor  string         `json:"backgroundColor"`
	PersonalGmail    string         `gorm:"type:varchar(100);index" json:"personal"`
	SupervisorGmail  *string        `gorm:"type:varchar(100);index" json:"supervisor"`
	Status           string         `gorm:"type:enum('requested', 'accepted', 'rejected' )" json:"status"`
	NilaiProgres     float64        `json:"nilaiProgres"`
	GoalStatus       string         `gorm:"type:enum('progres', 'completed')" json:"goalStatus"`
	RequestedAt      time.Time      `json:"requestAt"`
	AcceptedAt       time.Time      `json:"acceptAt"`
	RejectedAt       time.Time      `json:"rejectedAt"`
	IsActive         string         `gorm:"type:enum('1', '0')" json:"isActive"`
	List             []List         `gorm:"foreignKey:GoalID;references:ID" json:"List"`
	PersonalGoalCard []Card         `gorm:"foreignKey:GmailPersonal;references:PersonalGmail" json:"PersonalGoalCard"`
	Notification     []Notification `json:"notification"`
}

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
