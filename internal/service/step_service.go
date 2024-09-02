package service

import (
	data "email-sequence/internal/data"
	"email-sequence/internal/model"
)

type StepService interface {
	AddStep(sequenceID int, step *model.SequenceStep) error
	UpdateStep(sequenceID, stepID string, step *model.SequenceStep) (*model.SequenceStep, error)
	DeleteStep(sequenceID, stepID string) error
	GetStep(id, stepID string) (*model.SequenceStep, error)
}

type stepService struct {
	repo data.StepDataAccess
}

func NewStepService(repo data.StepDataAccess) StepService {
	return &stepService{repo}
}

func (s *stepService) AddStep(sequenceID int, step *model.SequenceStep) error {
	step.SequenceID = sequenceID
	return s.repo.CreateStep(step)
}

func (s *stepService) UpdateStep(sequenceID, stepID string, step *model.SequenceStep) (*model.SequenceStep, error) {
	return s.repo.UpdateStep(sequenceID, stepID, step)
}

func (s *stepService) DeleteStep(sequenceID, stepID string) error {
	return s.repo.DeleteStep(sequenceID, stepID)
}

func (s *stepService) GetStep(sequenceID, stepID string) (*model.SequenceStep, error) {
	step, err := s.repo.GetStep(sequenceID, stepID)
	if err != nil {
		return nil, err
	}
	return step, nil
}
