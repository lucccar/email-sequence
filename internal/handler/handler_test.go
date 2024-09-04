package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"email-sequence/internal/data"
	"email-sequence/internal/model"
	"email-sequence/internal/service"
)

var db *gorm.DB
var seqDataAccess data.SequenceDataAccess
var seqService service.SequenceService
var seqHandler *SequenceHandler

func setup() {
	var err error
	dsn := "host=localhost user=postgres password=secret dbname=mytestdatabase port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&model.Sequence{}, &model.SequenceStep{})
	if err != nil {
		panic("failed to migrate database")
	}

	seqDataAccess = data.NewSequenceDataAccess(db)
	seqService = service.NewSequenceService(seqDataAccess)
	seqHandler = NewSequenceHandler(seqService)
}

func dropDBs() {
	db.Exec("DROP TABLE IF EXISTS sequence_steps")
	db.Exec("DROP TABLE IF EXISTS sequences")
}

func TestCreateSequence(t *testing.T) {
	setup()
	defer dropDBs()

	r := gin.Default()
	r.POST("/sequences", seqHandler.CreateSequence)

	sequence := map[string]interface{}{
		"name":                   "Test Sequence",
		"open_tracking_enabled":  true,
		"click_tracking_enabled": true,
		"steps": []map[string]interface{}{
			{
				"subject": "Step 1 Subject",
				"content": "Step 1 Content",
			},
			{
				"subject": "Step 2 Subject",
				"content": "Step 2 Content",
			},
		},
	}

	body, _ := json.Marshal(sequence)
	req, _ := http.NewRequest("POST", "/sequences", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
	assert.Equal(t, "Test Sequence", response["name"])
	assert.Equal(t, 2, len(response["steps"].([]interface{})))
}

func TestGetSequence(t *testing.T) {
	setup()
	defer dropDBs()

	seq := model.Sequence{
		Name:                 "Test Sequence",
		OpenTrackingEnabled:  true,
		ClickTrackingEnabled: true,
		Steps: []model.SequenceStep{
			{Subject: "Step 1", Content: "Content 1", StepOrder: 1},
			{Subject: "Step 2", Content: "Content 2", StepOrder: 2},
		},
	}
	db.Create(&seq)

	r := gin.Default()
	r.GET("/sequences/:id", seqHandler.GetSequence)

	req, _ := http.NewRequest("GET", "/sequences/"+strconv.Itoa(seq.ID), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
	assert.Equal(t, "Test Sequence", response["name"])
	assert.Equal(t, 2, len(response["steps"].([]interface{})))

	var steps []model.SequenceStep

	if rawSteps, ok := response["steps"].([]interface{}); ok {
		mapstructure.Decode(rawSteps, &steps)
	}
	assert.Equal(t, seq.Steps[1].Subject, steps[1].Subject)
	assert.Equal(t, seq.Steps[1].Content, steps[1].Content)
}

func TestGetSequences(t *testing.T) {
	setup()
	defer dropDBs()

	seq1 := model.Sequence{
		Name:                 "Test Sequence 1",
		OpenTrackingEnabled:  true,
		ClickTrackingEnabled: true,
		Steps: []model.SequenceStep{
			{Subject: "Step 1", Content: "Content 1", StepOrder: 1},
		},
	}
	seq2 := model.Sequence{
		Name:                 "Test Sequence 2",
		OpenTrackingEnabled:  false,
		ClickTrackingEnabled: true,
		Steps: []model.SequenceStep{
			{Subject: "Step 2", Content: "Content 2", StepOrder: 1},
		},
	}
	db.Create(&seq1)
	db.Create(&seq2)

	r := gin.Default()
	r.GET("/sequences", seqHandler.GetSequences)

	req, _ := http.NewRequest("GET", "/sequences", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
	assert.Equal(t, 2, len(response))
}

func TestUpdateSequenceTracking(t *testing.T) {
	setup()
	defer dropDBs()

	sequence := model.Sequence{
		Name:                 "Test Sequence for Update Tracking 55",
		OpenTrackingEnabled:  false,
		ClickTrackingEnabled: false,
	}
	db.Create(&sequence)

	payload := map[string]bool{
		"open_tracking_enabled":  true,
		"click_tracking_enabled": true,
	}
	jsonPayload, err := json.Marshal(payload)
	assert.Nil(t, err)

	r := gin.Default()
	r.PATCH("/sequences/:id/tracking", seqHandler.UpdateSequenceTracking)

	req, _ := http.NewRequest("PATCH", "/sequences/"+strconv.Itoa(sequence.ID)+"/tracking", bytes.NewBuffer(jsonPayload))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedSequence model.Sequence
	db.First(&updatedSequence, sequence.ID)

	assert.True(t, updatedSequence.OpenTrackingEnabled)
	assert.True(t, updatedSequence.ClickTrackingEnabled)
}
