package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

// TimingResolver ..
type TimingResolver struct {
	T *model.Timing
}

// Id ..
func (r *TimingResolver) Id() string {
	return r.T.Id
}

// Event ..
func (r *TimingResolver) Event() *util.Time {
	return r.T.Event
}

// BoundsType ..
func (r *TimingResolver) BoundsType() *model.FrequenceBoundsType {
	return r.T.BoundsType
}

// BoundsDuration ..
func (r *TimingResolver) BoundsDuration() *int32 {
	return r.T.BoundsDuration
}

// BoundsRange ..
func (r *TimingResolver) BoundsRange() *RangeResolver {
	return &RangeResolver{r.T.BoundsRange}
}

// BoundsPeriod ..
func (r *TimingResolver) BoundsPeriod() *PeriodResolver {
	return &PeriodResolver{r.T.BoundsPeriod}
}

// Count ..
func (r *TimingResolver) Count() *int32 {
	return r.T.Count
}

// CountMax ..
func (r *TimingResolver) CountMax() *int32 {
	return r.T.CountMax
}

// Duration ..
func (r *TimingResolver) Duration() *float64 {
	return r.T.Duration
}

// DurationMax ..
func (r *TimingResolver) DurationMax() *float64 {
	return r.T.DurationMax
}

// DurationUnit ..
func (r *TimingResolver) DurationUnit() *model.UnitOfTime {
	return r.T.DurationUnit
}

// Frequency ..
func (r *TimingResolver) Frequency() *int32 {
	return r.T.Frequency
}

// FrequencyMax ..
func (r *TimingResolver) FrequencyMax() *int32 {
	return r.T.FrequencyMax
}

// Period ..
func (r *TimingResolver) Period() *float64 {
	return r.T.Period
}

// PeriodMax ..
func (r *TimingResolver) PeriodMax() *float64 {
	return r.T.PeriodMax
}

// PeriodUnit ..
func (r *TimingResolver) PeriodUnit() *model.UnitOfTime {
	return r.T.PeriodUnit
}

// DayOfWeek ..
func (r *TimingResolver) DayOfWeek() *model.DaysOfWeek {
	return r.T.DayOfWeek
}

// Time ..
func (r *TimingResolver) Time() *util.Time {
	return r.T.Time
}

// When ..
func (r *TimingResolver) When() *model.EventTiming {
	return r.T.When
}

// Offset ..
func (r *TimingResolver) Offset() *int32 {
	return r.T.Offset
}

// Code ..
func (r *TimingResolver) Code() *model.TimingAbbreviation {
	return r.T.Code
}
