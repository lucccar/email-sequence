package data

import (
	"email-sequence/internal/model"

	"gorm.io/gorm"
)

type SequenceDataAccess interface {
	CreateSequence(sequence *model.Sequence) (*model.Sequence, error)

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

func (r *sequenceDataAccess) CreateSequence(sequence *model.Sequence) (*model.Sequence, error) {
	result := r.db.Create(sequence)

	if result.Error != nil {
		return nil, result.Error
	}
	return sequence, nil
}

func (r *sequenceDataAccess) DeleteSequence(sequenceID int) error {

	if err := r.db.Where("sequence_id = ?", sequenceID).Delete(&model.SequenceStep{}).Error; err != nil {
		return err
	}

	return r.db.Delete(&model.Sequence{}, sequenceID).Error
}

func (r *sequenceDataAccess) GetSequence(sequenceID string) (*model.Sequence, error) {
	var sequence model.Sequence
	if err := r.db.Preload("Steps").First(&sequence, sequenceID).Error; err != nil {
		return nil, err
	}
	return &sequence, nil
}

func (r *sequenceDataAccess) GetSequences() ([]model.Sequence, error) {
	var sequences []model.Sequence

	if err := r.db.Preload("Steps").Find(&sequences).Error; err != nil {
		return nil, err
	}

	return sequences, nil
}

func (r *sequenceDataAccess) UpdateSequenceTracking(sequenceID string, openTracking, clickTracking bool) (*model.Sequence, error) {
	var sequence model.Sequence

	if err := r.db.Model(&sequence).Where("id = ?", sequenceID).
		Updates(map[string]interface{}{
			"open_tracking_enabled":  openTracking,
			"click_tracking_enabled": clickTracking,
		}).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("id = ?", sequenceID).First(&sequence).Error; err != nil {
		return nil, err
	}

	return &sequence, nil

}
