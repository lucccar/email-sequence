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
	if err := h.service.CreateSequence(&sequence); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sequence)
}

// Implement other handlers (UpdateSequence, DeleteSequence, etc.)
