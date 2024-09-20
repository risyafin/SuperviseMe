package entity

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	ID            int             `json:"id"`
	GmailPersonal string          `json:"gmailPersonal"`
	ListID        int             `json:"listId"`
	Cover         string          `json:"cover"`
	CardName      string          `json:"cardName"`
	Label         string          `json:"label"`
	Attachment    string          `json:"attachment"`
	Content       string          `json:"content"`
	Status        string          `gorm:"type:enum('progres', 'completed')" json:"status"`
	NilaiProgres  int             `json:"nilaiProgres"`
	StartDate     time.Time       `json:"startDate"` // 2024-09-28T00:00:00Z
	EndDate       time.Time       `json:"endDate"`
	CheckListCard []CheckListCard `json:"checkListCard"`
	Comment       []Comment       `json:"comment"`
}
type CheckListCard struct {
	gorm.Model
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CardID int    `json:"cardId"`
	Name   string `json:"name"`
	IsDone string `gorm:"type:enum('1', '0');default:'0'" json:"isDone"`
}

type CardResponHome struct {
	ID            int       `json:"id"`
	Status        string    `gorm:"type:enum('progres', 'completed')" json:"status"`
	NilaiProgres  int       `json:"nilaiProgres"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	GmailPersonal string    `json:"gmailPersonal"`
}

type ResponCreateCard struct {
	ID            int                     `json:"id"`
	CardName      string                  `json:"cardName"`
	Label         string                  `json:"label"`
	Attachment    string                  `json:"attachment"`
	Content       string                  `json:"content"`
	StartDate     time.Time               `json:"startDate"`
	EndDate       time.Time               `json:"endDate"`
	CheckListCard []CheckListCardResponse `json:"checkListCard"`
	ListID        int                     `json:"listId"`
}

type CardRespons struct {
	ID        int       `json:"id"`
	CardName  string    `json:"cardName"`
	Label     string    `json:"label"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	ListID    int       `json:"listId"`
}

type CheckListCardResponse struct {
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CardID int    `json:"cardId"`
	Name   string `json:"name"`
	IsDone string `gorm:"type:enum('1', '0')" json:"isDone"`
}