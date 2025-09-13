package model

import (
	"time"
)

//HealthKnowledgeGraphMessage represents deindefied consumer record in knowledge record
type HealthKnowledgeGraphMessage struct {
	Profiles *[]Profile               `json:"profiles"`
	Records  *[]HealthKnowledgeRecord `json:"records"`
}

// Profile ...
type Profile struct {
	ProfileID      string          `json:"profileID"`
	Gender         *Gender         `json:"gender,omitempty"`
	AgeGroup       *string         `json:"ageGroup,omitempty"`
	MarritalStatus *MarritalStatus `json:"marritalStatus,omitempty"`
	Country        *string         `json:"country,omitempty"`
	State          *string         `json:"state,omitempty"`
	Race           *Race           `json:"race,omitempty"`
	Ethnicity      *string         `json:"ethnicity,omitempty"`
}

// HealthKnowledgeRecord ...
type HealthKnowledgeRecord struct {
	ProfileID  string          `json:"profileID"`
	TimeStamp  *time.Time      `json:"timeStamp"`
	EntityType string          `json:"entityType"`
	Values     *[]Attribute    `json:"values"`
	Code       *CodableConcept `json:"codes"`
}
