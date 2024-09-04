package handler

import (
	"email-sequence/internal/model"
	"email-sequence/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SequenceHandler struct {
	service service.SequenceService
}

func NewSequenceHandler(s service.SequenceService) *SequenceHandler {
	return &SequenceHandler{service: s}
}

func (h *SequenceHandler) CreateSequence(c *gin.Context) {
	var sequence model.Sequence
	if err := c.ShouldBindJSON(&sequence); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i := range sequence.Steps {
		sequence.Steps[i].StepOrder = i + 1
	}
	insertedSequence, err := h.service.CreateSequence(&sequence)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, insertedSequence)
}

type SequenceTrackingUpdateInput struct {
	OpenTrackingEnabled  bool `json:"open_tracking_enabled"`
	ClickTrackingEnabled bool `json:"click_tracking_enabled"`
}

func (h *SequenceHandler) UpdateSequenceTracking(c *gin.Context) {
	sequenceID := c.Param("id")

	var update SequenceTrackingUpdateInput
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSequence, err := h.service.UpdateSequenceTracking(sequenceID, update.OpenTrackingEnabled, update.ClickTrackingEnabled)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedSequence)
}

func (h *SequenceHandler) GetSequence(c *gin.Context) {

	sequenceID := c.Param("id")

	sequence, err := h.service.GetSequence(sequenceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if sequence == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sequence not found"})
		return
	}

	c.JSON(http.StatusOK, sequence)
}

func (h *SequenceHandler) GetSequences(c *gin.Context) {
	sequences, err := h.service.GetSequences()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sequences)

}
