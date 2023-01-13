package service

import (
	"github.com/nailus/workout/internal/entity"
)

// TODO: разделить на разные сервисы
type Service struct {
	repository *Repository
}

func(s *Service) New(r *Repository) *Service {
	return &Service{repository: r}
}

func (s *Service) GetAllExercises() ([]entity.Exercise, error) {
	var exerciseList []entity.Exercise
	var exercise1 entity.Exercise
	exercise1.Id = 1
	exercise1.Title = "Title 1 Test"
	exercise1.Body = "Body 1 Test"

	var exercise2 entity.Exercise
	exercise2.Id = 2
	exercise2.Title = "Title 2 Test"
	exercise2.Body = "Body 2 Test"
	exerciseList = append(exerciseList, exercise1)
	exerciseList = append(exerciseList, exercise2)
	
	return exerciseList, nil
}