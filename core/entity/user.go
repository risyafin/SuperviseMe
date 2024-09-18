package entity

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type User struct {
	gorm.Model
	ID              int           `json:"id"`
	GoogleID        string        `gorm:"type:varchar(500);uniqueIndex" json:"googleId"`
	Name            string        `json:"name"`
	Email           string        `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Password        string        `json:"password"`
	Picture         string        `json:"picture"`
	Comment         []Comment     `json:"comment"`
	PersonalGoals   []Goals       `gorm:"foreignKey:PersonalGmail;references:Email" json:"personalGoals"`
	SupervisorGoals []Goals       `gorm:"foreignKey:SupervisorGmail;references:Email" json:"supervisorGoals"`
	ActivityLog     []ActivityLog `json:"activityLog"`
}

type UserResponProfile struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type UserResponseHome struct {
	ID              int                 `json:"id"`
	Name            string              `json:"name"`
	PersonalGoals   []GoalsResponseHome `gorm:"foreignKey:PersonalGmail;references:Email" json:"personalGoals"`
	SupervisorGoals []GoalsResponseHome `gorm:"foreignKey:SupervisorGmail;references:Email" json:"supervisorGoals"`
}

type UserResponseToken struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

type SupervisorRespons struct {
	Email           string                  `json:"email"`
	SupervisorGoals []GoalSupervisorRespons `json:"supervisor"`
}

type PersonalRespons struct {
	Email         string                `json:"email"`
	PersonalGoals []GoalPersonalRespons `json:"personal"`
}

type ActivityLog struct {
	gorm.Model
	ID     int    `json:"id"`
	Action string `json:"action"`
	UserID string `json:"userId"`
}
type ActivityLogRespon struct {
	ID     int    `json:"id"`
	Action string `json:"action"`
	UserID string `json:"userId"`
}

type MyClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	Id    int    `json:"id"`
}
