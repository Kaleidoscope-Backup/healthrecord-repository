package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
SleepStageResolver
================================*/

//SleepStageResolver ..
type SleepStageResolver struct {
	M *model.SleepStage
}

//Id ..
func (r *SleepStageResolver) Id() string {
	return r.M.Id
}

//Type ..
func (r *SleepStageResolver) Type() *model.SleepStageType {
	return r.M.Type
}

//Duration ..
func (r *SleepStageResolver) Duration() *float64 {
	return r.M.Duration
}

//Latency ..
func (r *SleepStageResolver) Latency() *float64 {
	return r.M.Latency
}

//TotalSleepTime ..
func (r *SleepStageResolver) TotalSleepTime() *float64 {
	return r.M.TotalSleepTime
}

//SleepPeriodTime ..
func (r *SleepStageResolver) SleepPeriodTime() *float64 {
	return r.M.SleepPeriodTime
}
