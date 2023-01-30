package service

import (
	"log"

	"github.com/nailus/workout/internal/entity"
	"github.com/nailus/workout/internal/repository"
	"golang.org/x/crypto/bcrypt"

	//"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// TODO: разделить на разные сервисы
type Service struct {
	repository *repository.Repository
}

type jwtTokenClaims struct {
	jwt.StandardClaims
	UserId int
}

const (
	signingKey = "vj4NgnfZG2PKhtGO"
	tokenTTL   = 1 * time.Hour
)

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

func (s *Service) GenerateAuthToken(user *entity.User) (string, error) {
	foundUser, _ := s.repository.GetUser(user.Email)

	if bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)) != nil {
		return "", nil
	}

	jwtTokenClaims := jwtTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		foundUser.Id,
	}

	

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtTokenClaims)


	return token.SignedString([]byte(signingKey))
}

