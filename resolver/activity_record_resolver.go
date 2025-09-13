package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
ActivityRecord Resolver
================================*/

//ActivityRecordResolver ..
type ActivityRecordResolver struct {
	HealthRecordResolver
	A *model.ActivityRecord
}

//Id ..
func (r *ActivityRecordResolver) Id() string {
	return r.A.Id
}

//ActivityType ..
func (r *ActivityRecordResolver) ActivityType() model.ActivityType {
	return r.A.ActivityType
}

//Frequency ..
func (r *ActivityRecordResolver) Frequency() *int32 {
	return r.A.Frequency
}

//FrequencyUnit ..
func (r *ActivityRecordResolver) FrequencyUnit() *string {
	return r.A.FrequencyUnit
}

//Duration ..
func (r *ActivityRecordResolver) Duration() *int32 {
	return r.A.Duration
}

//DurationUnit ..
func (r *ActivityRecordResolver) DurationUnit() *string {
	return r.A.DurationUnit
}

//Code ..
func (r *ActivityRecordResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.A.Code}
}

//Distance ..
func (r *ActivityRecordResolver) Distance() *int32 {
	return r.A.Distance
}

//DistanceUnit ..
func (r *ActivityRecordResolver) DistanceUnit() *string {
	return r.A.DistanceUnit
}

//Steps ..
func (r *ActivityRecordResolver) Steps() *int32 {
	return r.A.Steps
}

//Calories ..
func (r *ActivityRecordResolver) Calories() *float64 {
	return r.A.Calories
}

//CaloryUnit ..
func (r *ActivityRecordResolver) CaloryUnit() *string {
	return r.A.CaloryUnit
}

//Vigorous ..
func (r *ActivityRecordResolver) Vigorous() *int32 {
	return r.A.Vigorous
}

//Moderate ..
func (r *ActivityRecordResolver) Moderate() *int32 {
	return r.A.Moderate
}

//Light ..
func (r *ActivityRecordResolver) Light() *int32 {
	return r.A.Light
}

//Sedentary ..
func (r *ActivityRecordResolver) Sedentary() *int32 {
	return r.A.Sedentary
}
