package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

/*==============================
SleepRecordResolver
================================*/

// SleepRecordResolver ..
type SleepRecordResolver struct {
	HealthRecordResolver
	C *model.SleepRecord
}

// Id ..
func (r *SleepRecordResolver) Id() string {
	return r.C.Id
}

// StartTime ..
func (r *SleepRecordResolver) StartTime() util.Time {
	return r.C.StartTime
}

// EndTime ..
func (r *SleepRecordResolver) EndTime() util.Time {
	return r.C.EndTime
}

// MainSleep ..
func (r *SleepRecordResolver) MainSleep() *bool {
	return r.C.MainSleep
}

// TimeUnit ..
func (r *SleepRecordResolver) TimeUnit() *string {
	return r.C.TimeUnit
}

// TotalRecordingTime ..
func (r *SleepRecordResolver) TotalRecordingTime() *float64 {
	return r.C.TotalRecordingTime
}

// TotalSleepTime ..
func (r *SleepRecordResolver) TotalSleepTime() *float64 {
	return r.C.TotalSleepTime
}

// TimeAwake ..
func (r *SleepRecordResolver) TimeAwake() *float64 {
	return r.C.TimeAwake
}

// SleepEfficiency ..
func (r *SleepRecordResolver) SleepEfficiency() *float64 {
	return r.C.SleepEfficiency
}

// TimeToFallAsleep ..
func (r *SleepRecordResolver) TimeToFallAsleep() *float64 {
	return r.C.TimeToFallAsleep
}

// NumberOfAwekenings ..
func (r *SleepRecordResolver) NumberOfAwekenings() *int32 {
	return r.C.NumberOfAwekenings
}

// TimeAfterWakeup ..
func (r *SleepRecordResolver) TimeAfterWakeup() *float64 {
	return r.C.TimeAfterWakeup
}

// TimeInBed ..
func (r *SleepRecordResolver) TimeInBed() *float64 {
	return r.C.TimeInBed
}

// SleepStageSummary array ..
func (r *SleepRecordResolver) SleepStageSummary() *[]*SleepStageResolver {

	if r.C.SleepStageSummary != nil {
		var crs []*SleepStageResolver
		var cs []model.SleepStage
		cs = *r.C.SleepStageSummary

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.SleepStage
				c = cs[i]
				if cr := resolveSleepStageResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolveSleepStageResolver(c *model.SleepStage) *SleepStageResolver {
	return &SleepStageResolver{c}
}

// SleepContinuities array ..
func (r *SleepRecordResolver) SleepContinuities() *[]*SleepContinuityResolver {

	if r.C.SleepContinuities != nil {
		var crs []*SleepContinuityResolver
		var cs []model.SleepContinuity
		cs = *r.C.SleepContinuities

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.SleepContinuity
				c = cs[i]
				if cr := resolveSleepContinuityResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolveSleepContinuityResolver(c *model.SleepContinuity) *SleepContinuityResolver {
	return &SleepContinuityResolver{c}
}
