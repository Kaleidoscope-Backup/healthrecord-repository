package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
ClinicalAssesmentObservation Resolver
================================*/

// ClinicalAssesmentObservationResolver ..
type ClinicalAssesmentObservationResolver struct {
	C *model.ClinicalAssesmentObservation
}

// Id ..
func (r *ClinicalAssesmentObservationResolver) Id() string {
	return r.C.Id
}

// Name ..
func (r *ClinicalAssesmentObservationResolver) Name() string {
	return r.C.Name
}

// Value ..
func (r *ClinicalAssesmentObservationResolver) Value() string {
	return r.C.Value
}

// Score ..
func (r *ClinicalAssesmentObservationResolver) Score() *int32 {
	return r.C.Score
}

// Code ..
func (r *ClinicalAssesmentObservationResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.C.Code}
}
