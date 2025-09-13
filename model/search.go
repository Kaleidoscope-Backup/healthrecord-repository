package model

import "github.com/Kaleidoscope-Backup/healthrecord-repository/util"

// PageInfo ...
type PageInfo struct {
	StartCursor *string `json:"startCursor"`
	EndCursor   *string `json:"endCursor"`
	HasNext     bool    `json:"hasNext" bson:"qualification"`
}

// SearchInput ...
type SearchInput struct {
	ConsumerID    string              `json:"consumerID"`
	Filters       *[]HealthRecordType `json:"filters"`
	Limit         *int32              `json:"limit"`
	LastEndCursor *string             `json:"lastEndCursor"`
	Start         *util.Time          `json:"start"`
	End           *util.Time          `json:"end"`
}
