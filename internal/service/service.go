package service

import (
	"github.com/nailus/workout/internal/entity"
	"github.com/nailus/workout/internal/repository"
)

// TODO: разделить на разные сервисы
type Service struct {
	repository *repository.Repository
}

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