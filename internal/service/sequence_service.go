package service

import (
	data "email-sequence/internal/data"
	"email-sequence/internal/model"
	"fmt"
)

type SequenceService interface {
	CreateSequence(sequence *model.Sequence) error
	UpdateSequence(sequence *model.Sequence) error
	GetSequence(sequenceID string) (*model.Sequence, error)
	GetSequences() ([]model.Sequence, error)
	UpdateSequenceTracking(sequenceID string, openTracking, clickTracking bool) (*model.Sequence, error)
}

type sequenceService struct {
	repo data.SequenceDataAccess
}

func NewSequenceService(repo data.SequenceDataAccess) SequenceService {
	return &sequenceService{repo}
}

// CreateSequence creates a new sequence with steps
func (s *sequenceService) CreateSequence(sequence *model.Sequence) error {
	fmt.Println(sequence)
	return s.repo.CreateSequence(sequence)
}

func (s *sequenceService) UpdateSequence(sequence *model.Sequence) error {
	return s.repo.UpdateSequence(sequence)
}

func (s *sequenceService) UpdateSequenceTracking(sequenceID string, openTracking, clickTracking bool) (*model.Sequence, error) {
	return s.repo.UpdateSequenceTracking(sequenceID, openTracking, clickTracking)

}

func (s *sequenceService) GetSequence(sequenceID string) (*model.Sequence, error) {
	sequence, err := s.repo.GetSequence(sequenceID)
	if err != nil {
		return nil, err
	}
	return sequence, nil
}

func (s *sequenceService) GetSequences() ([]model.Sequence, error) {
	sequences, err := s.repo.GetSequences()
	if err != nil {
		return nil, err
	}
	return sequences, nil
}
