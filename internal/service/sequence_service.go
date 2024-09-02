package service

import (
	data "email-sequence/internal/data"
	"email-sequence/internal/model"
)

type SequenceService interface {
	CreateSequence(sequence *model.Sequence) error
	UpdateSequence(sequence *model.Sequence) error
	// DeleteSequence(sequenceID int) error
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
	return s.repo.CreateSequence(sequence)
}

func (s *sequenceService) UpdateSequence(sequence *model.Sequence) error {
	return s.repo.UpdateSequence(sequence)
}

func (s *sequenceService) UpdateSequenceTracking(sequenceID string, openTracking, clickTracking bool) (*model.Sequence, error) {
	return s.repo.UpdateSequenceTracking(sequenceID, openTracking, clickTracking)

}
