package data

import (
	"email-sequence/internal/model"
	"errors"

	"gorm.io/gorm"
)

type StepDataAccess interface {
	CreateStep(step *model.SequenceStep) error
	UpdateStep(sequenceID, stepID string, step *model.SequenceStep) (*model.SequenceStep, error)
	DeleteStep(sequenceID, stepID string) error
	GetStep(sequenceID, stepID string) (*model.SequenceStep, error)
}

type stepDataAccess struct {
	db *gorm.DB
}

func NewStepDataAccess(db *gorm.DB) StepDataAccess {
	return &stepDataAccess{db}
}

func (r *stepDataAccess) CreateStep(step *model.SequenceStep) error {
	return r.db.Create(step).Error
}

func (r *stepDataAccess) UpdateStep(sequenceID, stepID string, step *model.SequenceStep) (*model.SequenceStep, error) {

	var existingStep model.SequenceStep
	if err := r.db.Where("id = ? AND sequence_id = ?", stepID, sequenceID).First(&existingStep).Error; err != nil {
		return nil, err
	}

	existingStep.Subject = step.Subject
	existingStep.Content = step.Content

	if err := r.db.Save(&existingStep).Error; err != nil {
		return nil, err
	}

	return &existingStep, nil
}

func (r *stepDataAccess) DeleteStep(sequenceID, stepID string) error {
	query := `DELETE FROM sequence_steps WHERE id=$1 AND sequence_id=$2`
	return r.db.Exec(query, stepID, sequenceID).Error
}

func (r *stepDataAccess) GetStep(sequenceID, stepID string) (*model.SequenceStep, error) {
	var step model.SequenceStep

	if err := r.db.Where("id = ? AND sequence_id = ?", stepID, sequenceID).First(&step).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil if no record is found
		}
		return nil, err // Return error if something goes wrong
	}

	return &step, nil

}
