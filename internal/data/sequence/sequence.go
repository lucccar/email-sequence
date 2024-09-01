package sequence

import (
	"email-sequence/internal/model"

	"gorm.io/gorm"
)

type SequenceDataAccess interface {
	CreateSequence(sequence *model.Sequence) error
	UpdateSequence(sequence *model.Sequence) error
	DeleteSequence(sequenceID int) error
	GetSequence(sequenceID int) (*model.Sequence, error)
	CreateStep(step *model.SequenceStep) error
	UpdateStep(step *model.SequenceStep) error
	DeleteStep(stepID int) error
}

type sequenceDataAccess struct {
	db *gorm.DB
}

func NewSequenceDataAccess(db *gorm.DB) SequenceDataAccess {
	return &sequenceDataAccess{db}
}

// CreateSequence inserts a new sequence into the database
func (r *sequenceDataAccess) CreateSequence(sequence *model.Sequence) error {
	return r.db.Create(sequence).Error
}

// UpdateSequence updates an existing sequence in the database
func (r *sequenceDataAccess) UpdateSequence(sequence *model.Sequence) error {
	return r.db.Save(sequence).Error
}

// DeleteSequence deletes a sequence and its associated steps from the database
func (r *sequenceDataAccess) DeleteSequence(sequenceID int) error {
	// First delete associated steps
	if err := r.db.Where("sequence_id = ?", sequenceID).Delete(&model.SequenceStep{}).Error; err != nil {
		return err
	}
	// Then delete the sequence
	return r.db.Delete(&model.Sequence{}, sequenceID).Error
}

// GetSequence retrieves a sequence and its associated steps from the database
func (r *sequenceDataAccess) GetSequence(sequenceID int) (*model.Sequence, error) {
	var sequence model.Sequence
	if err := r.db.Preload("Steps").First(&sequence, sequenceID).Error; err != nil {
		return nil, err
	}
	return &sequence, nil
}

// CreateStep adds a new step to an existing sequence
func (r *sequenceDataAccess) CreateStep(step *model.SequenceStep) error {
	return r.db.Create(step).Error
}

// UpdateStep updates an existing step in a sequence
func (r *sequenceDataAccess) UpdateStep(step *model.SequenceStep) error {
	return r.db.Save(step).Error
}

// DeleteStep deletes a step from a sequence
func (r *sequenceDataAccess) DeleteStep(stepID int) error {
	return r.db.Delete(&model.SequenceStep{}, stepID).Error
}
