package mocks

import (
	"email-sequence/internal/service" // Adjust the import path to your project structure

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// MockSequenceService is a mock implementation of the SequenceService interface
type MockSequenceService struct {
	mock.Mock
}

// CreateSequence is a mock method for the CreateSequence function
func (m *MockSequenceService) CreateSequence(ctx *gin.Context, sequence service.SequenceService) error {
	args := m.Called(ctx, sequence)
	return args.Error(0)
}

// GetSequence is a mock method for the GetSequence function
func (m *MockSequenceService) GetSequence(ctx *gin.Context, id int) (service.SequenceService, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(service.SequenceService), args.Error(1)
}

// UpdateSequenceTracking is a mock method for the UpdateSequenceTracking function
func (m *MockSequenceService) UpdateSequenceTracking(ctx *gin.Context, id int, openTrackingEnabled bool, clickTrackingEnabled bool) error {
	args := m.Called(ctx, id, openTrackingEnabled, clickTrackingEnabled)
	return args.Error(0)
}

// GetSequences is a mock method for the GetSequences function
func (m *MockSequenceService) GetSequences(ctx *gin.Context) ([]service.SequenceService, error) {
	args := m.Called(ctx)
	return args.Get(0).([]service.SequenceService), args.Error(1)
}
