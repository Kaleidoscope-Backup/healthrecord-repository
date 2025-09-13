package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
ObservationReecord Resolver
================================*/

// ObservationRecordResolver ..
type ObservationRecordResolver struct {
	HealthRecordResolver
	O *model.ObservationRecord
}

// Id ..
func (r *ObservationRecordResolver) Id() string {
	return r.O.Id
}

// Status ..
func (r *ObservationRecordResolver) Status() model.ObservationStatus {
	return r.O.Status
}

// Category ..
func (r *ObservationRecordResolver) Category() model.ObservationCategory {
	return r.O.Category
}

// CategoryCode ..
func (r *ObservationRecordResolver) CategoryCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.O.CategoryCode}
}

// Code ..
func (r *ObservationRecordResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.O.Code}
}

// Performer ..
func (r *ObservationRecordResolver) Performer() *ReferenceActorResolver {
	return &ReferenceActorResolver{r.O.Performer}
}

// Effective ..
func (r *ObservationRecordResolver) Effective() *PeriodResolver {
	return &PeriodResolver{r.O.Effective}
}

// Value ..
func (r *ObservationRecordResolver) Value() *ValueResolver {
	return &ValueResolver{&r.O.Value}
}

// DataAbsentReason ..
func (r *ObservationRecordResolver) DataAbsentReason() *string {
	return r.O.DataAbsentReason
}

// DataAbsentReasonCode ..
func (r *ObservationRecordResolver) DataAbsentReasonCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.O.DataAbsentReasonCode}
}

// Interpretation ..
func (r *ObservationRecordResolver) Interpretation() *string {
	return r.O.Interpretation
}

// InterpretationCode ..
func (r *ObservationRecordResolver) InterpretationCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.O.InterpretationCode}
}

// Comment ..
func (r *ObservationRecordResolver) Comment() *string {
	return r.O.Comment
}

// BodySite ..
func (r *ObservationRecordResolver) BodySite() *string {
	return r.O.BodySite
}

// BodySiteCode ..
func (r *ObservationRecordResolver) BodySiteCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.O.BodySiteCode}
}

// Method ..
func (r *ObservationRecordResolver) Method() *string {
	return r.O.Method
}

// MethodCode ..
func (r *ObservationRecordResolver) MethodCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.O.MethodCode}
}

// Device ..
func (r *ObservationRecordResolver) Device() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{r.O.Device}
}

// ReferenceRange array ..
func (r *ObservationRecordResolver) ReferenceRange() *[]*ReferenceRangeResolver {

	if r.O.ReferenceRange != nil {
		var crs []*ReferenceRangeResolver
		var cs []model.ReferenceRange
		cs = *r.O.ReferenceRange

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ReferenceRange
				c = cs[i]
				if cr := ResolveReferenceRangeResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
