package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

//SymptomResolver ..
type SymptomResolver struct {
	M *model.Symptom
}

//Id ..
func (r *SymptomResolver) Id() string {
	return r.M.Id
}

//Name ..
func (r *SymptomResolver) Name() string {
	return r.M.Name
}

//Code ..
func (r *SymptomResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.M.Code}
}
