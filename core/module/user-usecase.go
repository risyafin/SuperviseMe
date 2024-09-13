package module

import (
	"errors"
	"fmt"
	"math/rand"
	"superviseMe/core/entity"
	"superviseMe/core/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase interface {
	GetGoalsBySuperviseeUser(email string) (*entity.User, error)
	GetGoalSupervisor(email string) (*entity.User, error)
	GetGoalPersonal(email string) (*entity.User, error)
	Login(email string, password string) (string, *entity.User, error)
	SaveUser(user *entity.User) error
	GetUserByID(id string) (*entity.User, error)
	Registration(user *entity.User) error
	UpdateName(name string, email string) error
}

type userUseCase struct {
	userRepository repository.UserResitory
}

func NewUserUseCase(userRepository repository.UserResitory) UserUseCase {
	return &userUseCase{userRepository: userRepository}
}

func (e *userUseCase) GetGoalsBySuperviseeUser(email string) (*entity.User, error) {
	return e.userRepository.GetGoalsBySuperviseeUser(email)
}

func (e *userUseCase) GetGoalSupervisor(email string) (*entity.User, error) {
	return e.userRepository.GetGoalSupervisor(email)
}

func (e *userUseCase) GetGoalPersonal(email string) (*entity.User, error) {
	return e.userRepository.GetGoalSupervisor(email)
}

func (u *userUseCase) UpdateName(name string, email string) error {
	fmt.Println("usecase", email, "&", name)
	return u.userRepository.UpdateName(name, email)
}

func (u *userUseCase) SaveUser(user *entity.User) error {
	return u.userRepository.Save(user)
}

func (u userUseCase) GetUserByID(id string) (*entity.User, error) {
	return u.userRepository.FindByID(id)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}
func validatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	return nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func SeededRandString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (e *userUseCase) Registration(user *entity.User) error {
	err := validatePassword(user.Password)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	} else {
		hash, _ := HashPassword(user.Password)
		user.Password = hash
		user.Name = SeededRandString(10)
		errs := e.userRepository.Registration(user)
		return errs
	}
}
func (e *userUseCase) Login(email string, password string) (string, *entity.User, error) {
	fmt.Println("ini di usecase", email)
	user, err := e.userRepository.GetUserByGmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", user, errors.New("invalid idential")
		}
		return "", user, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", user, errors.New("invalid credentials")
	}

	claims := entity.MyClaims{
		Email:          user.Email,
		Id:             user.ID,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err1 := token.SignedString([]byte("Bolong"))
	if err != nil {
		return "", user, err1
	}

	return signedToken, user, nil
}
