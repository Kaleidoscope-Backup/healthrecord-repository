package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
AdverseEventRecord Resolver
================================*/

//AdverseEventRecordResolver ..
type AdverseEventRecordResolver struct {
	HealthRecordResolver
	A *model.AdverseEventRecord
}

//Id ..
func (r *AdverseEventRecordResolver) Id() string {
	return r.A.Id
}

//Category ..
func (r *AdverseEventRecordResolver) Category() model.AdverseEventCategory {
	return r.A.Category
}

//CategoryCode ..
func (r *AdverseEventRecordResolver) CategoryCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.A.CategoryCode}
}

//EventType ..
func (r *AdverseEventRecordResolver) EventType() string {
	return r.A.EventType
}

//EventTypeCode ..
func (r *AdverseEventRecordResolver) EventTypeCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.A.EventTypeCode}
}

//Location ..
func (r *AdverseEventRecordResolver) Location() *GeoLocationResolver {
	return &GeoLocationResolver{r.A.Location}
}

//Seriousness ..
func (r *AdverseEventRecordResolver) Seriousness() *model.Severity {
	return r.A.Seriousness
}

//Outcome ..
func (r *AdverseEventRecordResolver) Outcome() *string {
	return r.A.Outcome
}

//OutcomeCode ..
func (r *AdverseEventRecordResolver) OutcomeCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.A.OutcomeCode}
}

//Recorder ..
func (r *AdverseEventRecordResolver) Recorder() *ReferenceActorResolver {
	return &ReferenceActorResolver{r.A.Recorder}
}
