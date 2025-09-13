package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
LabResultObservationRecord Resolver
================================*/

//LabResultObservationRecordResolver ..
type LabResultObservationRecordResolver struct {
	HealthRecordResolver
	C *model.LabResultObservationRecord
}

//Id ..
func (r *LabResultObservationRecordResolver) Id() string {
	return r.C.Id
}

//Category ..
func (r *LabResultObservationRecordResolver) Category() string {
	return r.C.Category
}

//Code ..
func (r *LabResultObservationRecordResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.C.Code}
}

//Specimen ..
func (r *LabResultObservationRecordResolver) Specimen() *string {
	return r.C.Specimen
}

//Comment ..
func (r *LabResultObservationRecordResolver) Comment() *string {
	return r.C.Comment
}

//Method ..
func (r *LabResultObservationRecordResolver) Method() *string {
	return r.C.Method
}

//Interpretation ..
func (r *LabResultObservationRecordResolver) Interpretation() *string {
	return r.C.Interpretation
}

//MethodCode ..
func (r *LabResultObservationRecordResolver) MethodCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.C.MethodCode}
}

//Observations array ..
func (r *LabResultObservationRecordResolver) Observations() *[]*LabResultObservationResolver {

	if r.C.Observations != nil {
		var crs []*LabResultObservationResolver
		var cs []model.LabResultObservation
		cs = *r.C.Observations

		if r.C.Observations != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.LabResultObservation
				c = cs[i]
				if cr := resolveLabResultObservation(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolveLabResultObservation(c *model.LabResultObservation) *LabResultObservationResolver {
	return &LabResultObservationResolver{c}
}
