package service

import (
	"email-sequence/internal/data/sequence"
	"email-sequence/internal/model"
)

type SequenceService interface {
	CreateSequence(sequence *model.Sequence) error
	UpdateSequence(sequence *model.Sequence) error
	DeleteSequence(sequenceID int) error
	AddStep(sequenceID int, step *model.SequenceStep) error
	UpdateStep(step *model.SequenceStep) error
	DeleteStep(stepID int) error
	UpdateTracking(sequenceID int, openTracking, clickTracking bool) error
}

type sequenceService struct {
	repo sequence.SequenceDataAccess
}

func NewSequenceService(repo sequence.SequenceDataAccess) SequenceService {
	return &sequenceService{repo}
}

// CreateSequence creates a new sequence with steps
func (s *sequenceService) CreateSequence(sequence *model.Sequence) error {
	return s.repo.CreateSequence(sequence)
}

// UpdateSequence updates an existing sequence, including its steps
func (s *sequenceService) UpdateSequence(sequence *model.Sequence) error {
	return s.repo.UpdateSequence(sequence)
}

// DeleteSequence deletes a sequence and its steps
func (s *sequenceService) DeleteSequence(sequenceID int) error {
	return s.repo.DeleteSequence(sequenceID)
}

// AddStep adds a new step to an existing sequence
func (s *sequenceService) AddStep(sequenceID int, step *model.SequenceStep) error {
	// Ensure the step is associated with the correct sequence
	step.SequenceID = sequenceID
	return s.repo.CreateStep(step)
}

// UpdateStep updates an existing step in a sequence
func (s *sequenceService) UpdateStep(step *model.SequenceStep) error {
	return s.repo.UpdateStep(step)
}

// DeleteStep deletes a step from a sequence
func (s *sequenceService) DeleteStep(stepID int) error {
	return s.repo.DeleteStep(stepID)
}

// UpdateTracking updates the open and click tracking settings for a sequence
func (s *sequenceService) UpdateTracking(sequenceID int, openTracking, clickTracking bool) error {
	// Retrieve the sequence to update its tracking settings
	sequence, err := s.repo.GetSequence(sequenceID)
	if err != nil {
		return err
	}

	sequence.OpenTrackingEnabled = openTracking
	sequence.ClickTrackingEnabled = clickTracking

	return s.repo.UpdateSequence(sequence)
}
