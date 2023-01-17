package service

import (
	"github.com/nailus/workout/internal/entity"
	"github.com/nailus/workout/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	//"github.com/golang-jwt/jwt/v4"

)

// TODO: разделить на разные сервисы
type Service struct {
	repository *repository.Repository
}

// type Claims struct {
// 	UserId int
// 	jwt.StandardClaims
// }

func New(r *repository.Repository) *Service {
	return &Service{repository: r}
}

func (s *Service) GetAllExercises() ([]entity.Exercise, error) {
	exerciseList, err := s.repository.GetAllExercises()
	if err != nil {
		return nil, err
	}
	return exerciseList, nil
}

func (s *Service) CreateUser(user *entity.User) (int, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		log.Fatal(err)
		//return 0, err
	}
	user.Password = string(password)
	
	return s.repository.CreateUser(user)
}

func (s *Service) GetUser(user *entity.User) (*entity.User, error) {
	encryptPassword(user)
	return s.repository.GetUser(user)
} 

func encryptPassword(user *entity.User) {
	entryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(entryptedPassword)
}

