package handler

import (
	"email-sequence/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StepHandler struct {
	service service.StepService
}

type UpdateStepInput struct {
	Subject *string `json:"subject,omitempty"`
	Content *string `json:"content,omitempty"`
}

func NewStepHandler(s service.StepService) *StepHandler {
	return &StepHandler{service: s}
}

func (h *StepHandler) UpdateStep(c *gin.Context) {

	sequenceID := c.Param("id")
	stepID := c.Param("stepId")

	var input UpdateStepInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	step, err := h.service.GetStep(sequenceID, stepID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if input.Subject != nil {
		step.Subject = *input.Subject
	}
	if input.Content != nil {
		step.Content = *input.Content
	}

	updatedStep, err := h.service.UpdateStep(sequenceID, stepID, step)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedStep)

}

func (h *StepHandler) DeleteStep(c *gin.Context) {
	sequenceID := c.Param("id")
	stepID := c.Param("stepId")

	err := h.service.DeleteStep(sequenceID, stepID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
