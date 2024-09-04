package service_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"email-sequence/internal/data/mocks"
	"email-sequence/internal/model"
	"email-sequence/internal/service"
)

func TestCreateSequence(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.SequenceRepository)
	seqService := service.NewSequenceService(mockRepo)

	sequence := &model.Sequence{
		Name:                 "Test Sequence",
		OpenTrackingEnabled:  true,
		ClickTrackingEnabled: true,
		Steps: []model.SequenceStep{
			{Subject: "Step 1", Content: "Content for Step 1", StepOrder: 1},
			{Subject: "Step 2", Content: "Content for Step 2", StepOrder: 2},
		},
	}

	// Define what the mock should expect and return
	mockRepo.On("CreateSequence", sequence).Return(nil)

	// Act
	err := seqService.CreateSequence(sequence)

	// Assert
	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
