package model

type Sequence struct {
	ID                   int            `json:"id"`
	Name                 string         `json:"name"`
	OpenTrackingEnabled  bool           `json:"open_tracking_enabled"`
	ClickTrackingEnabled bool           `json:"click_tracking_enabled"`
	WaitHours            int            `json:"wait_hours"`
	Steps                []SequenceStep `json:"steps" gorm:"foreignKey:SequenceID"`
}

type SequenceStep struct {
	ID         int    `json:"id"`
	SequenceID int    `json:"sequence_id"`
	Subject    string `json:"subject"`
	Content    string `json:"content"`
	StepOrder  int    `json:"step_order"`
}
