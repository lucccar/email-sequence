package data

import (
	"email-sequence/internal/model"

	"gorm.io/gorm"
)

type SequenceDataAccess interface {
	CreateSequence(sequence *model.Sequence) error
	UpdateSequence(sequence *model.Sequence) error
	// DeleteSequence(sequenceID string) error
	GetSequence(sequenceID string) (*model.Sequence, error)
	GetSequences() ([]model.Sequence, error)
	UpdateSequenceTracking(sequenceID string, openTracking, clickTracking bool) (*model.Sequence, error)
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
func (r *sequenceDataAccess) GetSequence(sequenceID string) (*model.Sequence, error) {
	var sequence model.Sequence
	if err := r.db.Preload("Steps").First(&sequence, sequenceID).Error; err != nil {
		return nil, err
	}
	return &sequence, nil
}

func (r *sequenceDataAccess) GetSequences() ([]model.Sequence, error) {
	var sequences []model.Sequence

	// Query the database to find all sequences and preload the related steps
	if err := r.db.Preload("Steps").Find(&sequences).Error; err != nil {
		return nil, err
	}

	return sequences, nil
}

// UpdateSequenceTracking updates the tracking settings for a sequence in the database
func (r *sequenceDataAccess) UpdateSequenceTracking(sequenceID string, openTracking, clickTracking bool) (*model.Sequence, error) {
	var sequence model.Sequence

	// Update the sequence tracking fields
	if err := r.db.Model(&sequence).Where("id = ?", sequenceID).
		Updates(map[string]interface{}{
			"open_tracking_enabled":  openTracking,
			"click_tracking_enabled": clickTracking,
		}).Error; err != nil {
		return nil, err
	}

	// Retrieve the updated sequence
	if err := r.db.Where("id = ?", sequenceID).First(&sequence).Error; err != nil {
		return nil, err
	}

	return &sequence, nil

}
