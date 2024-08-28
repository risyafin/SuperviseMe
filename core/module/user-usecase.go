package module

import (
	"errors"
	"fmt"
	"superviseMe/core/entity"
	"superviseMe/core/repository"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase interface {
	GetGoalsBySuperviseeUser(gmail string) (*entity.User, error)
	GetGoalSupervisor(gmail string) (*entity.User, error)
	GetGoalPersonal(gmail string) (*entity.User, error)
	Login(gmail string, password string) (string, *entity.User, error)
	SaveUser(user *entity.User) error
	GetUserByID(id string) (*entity.User, error)
	Registration(user *entity.User) error
}

type userUseCase struct {
	userRepository repository.UserResitory
}

func NewUserUseCase(userRepository repository.UserResitory) UserUseCase {
	return &userUseCase{userRepository: userRepository}
}

func (e *userUseCase) GetGoalsBySuperviseeUser(gmail string) (*entity.User, error) {
	return e.userRepository.GetGoalsBySuperviseeUser(gmail)
}

func (e *userUseCase) GetGoalSupervisor(gmail string) (*entity.User, error) {
	return e.userRepository.GetGoalSupervisor(gmail)
}

func (e *userUseCase) GetGoalPersonal(gmail string) (*entity.User, error) {
	return e.userRepository.GetGoalSupervisor(gmail)
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

func (e *userUseCase) Registration(user *entity.User) error {
	err := validatePassword(user.Password)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	} else {
		hash, _ := HashPassword(user.Password)
		user.Password = hash
		errs := e.userRepository.Registration(user)
		return errs
	}
}
func (e *userUseCase) Login(gmail string, password string) (string, *entity.User, error) {
	fmt.Println("ini di usecase", gmail)
	user, err := e.userRepository.GetUserByGmail(gmail)
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
