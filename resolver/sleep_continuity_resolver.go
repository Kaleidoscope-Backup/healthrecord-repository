package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
SleepContinuityResolver
================================*/

// SleepContinuityResolver ..
type SleepContinuityResolver struct {
	M *model.SleepContinuity
}

// Id ..
func (r *SleepContinuityResolver) Id() string {
	return r.M.Id
}

// SourceOfArousal ..
func (r *SleepContinuityResolver) SourceOfArousal() *string {
	return r.M.SourceOfArousal
}

// NermCount ..
func (r *SleepContinuityResolver) NermCount() *int32 {
	return r.M.NermCount
}

// NermIndex ..
func (r *SleepContinuityResolver) NermIndex() *float64 {
	return r.M.NermIndex
}

// RemCount ..
func (r *SleepContinuityResolver) RemCount() *int32 {
	return r.M.RemCount
}

// RemIndex ..
func (r *SleepContinuityResolver) RemIndex() *float64 {
	return r.M.RemIndex
}

// TotalCount ..
func (r *SleepContinuityResolver) TotalCount() *int32 {
	return r.M.TotalCount
}

// TotalIndex ..
func (r *SleepContinuityResolver) TotalIndex() *float64 {
	return r.M.TotalIndex
}
