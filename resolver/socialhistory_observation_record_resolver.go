package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
HabitRecord Resolver
================================*/

// SocialHistoryObservationRecordResolver ..
type SocialHistoryObservationRecordResolver struct {
	HealthRecordResolver
	U *model.SocialHistoryObservationRecord
}

// Id ..
func (r *SocialHistoryObservationRecordResolver) Id() string {
	return r.U.Id
}

// Type ..
func (r *SocialHistoryObservationRecordResolver) Type() model.SocialHistoryObservationType {
	return r.U.Type
}

// Status ..
func (r *SocialHistoryObservationRecordResolver) Status() model.SocialHistoryObservationStatus {
	return r.U.Status
}

// Duration ..
func (r *SocialHistoryObservationRecordResolver) Duration() *int32 {
	return r.U.Duration
}

// DurationUnit ..
func (r *SocialHistoryObservationRecordResolver) DurationUnit() *string {
	return r.U.DurationUnit
}

// Start ..
func (r *SocialHistoryObservationRecordResolver) Start() *util.Time {
	return r.U.Start
}

// End ..
func (r *SocialHistoryObservationRecordResolver) End() *util.Time {
	return r.U.End
}

// Code ..
func (r *SocialHistoryObservationRecordResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.U.Code}
}

// Value ..
func (r *SocialHistoryObservationRecordResolver) Value() *ValueResolver {
	return &ValueResolver{r.U.Value}
}
