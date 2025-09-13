package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
ClinicalAssesmentObservationRecord Resolver
================================*/

//ClinicalAssesmentObservationRecordResolver ..
type ClinicalAssesmentObservationRecordResolver struct {
	HealthRecordResolver
	C *model.ClinicalAssesmentObservationRecord
}

//Id ..
func (r *ClinicalAssesmentObservationRecordResolver) Id() string {
	return r.C.Id
}

//Code ..
func (r *ClinicalAssesmentObservationRecordResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.C.Code}
}

//Comment ..
func (r *ClinicalAssesmentObservationRecordResolver) Comment() *string {
	return r.C.Comment
}

//Method ..
func (r *ClinicalAssesmentObservationRecordResolver) Method() *string {
	return r.C.Method
}

//MethodCode ..
func (r *ClinicalAssesmentObservationRecordResolver) MethodCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.C.MethodCode}
}

//Interpretation ..
func (r *ClinicalAssesmentObservationRecordResolver) Interpretation() *string {
	return r.C.Interpretation
}

//Observations array ..
func (r *ClinicalAssesmentObservationRecordResolver) Observations() *[]*ClinicalAssesmentObservationResolver {

	if r.C.Observations != nil {
		var crs []*ClinicalAssesmentObservationResolver
		var cs []model.ClinicalAssesmentObservation
		cs = *r.C.Observations

		if r.C.Observations != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ClinicalAssesmentObservation
				c = cs[i]
				if cr := resolveClinicalAssesmentObservation(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolveClinicalAssesmentObservation(c *model.ClinicalAssesmentObservation) *ClinicalAssesmentObservationResolver {
	return &ClinicalAssesmentObservationResolver{c}
}
