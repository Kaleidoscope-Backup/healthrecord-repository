package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
PersonalCharacteristicsObservation Resolver
================================*/

//PersonalCharacteristicsObservationResolver ..
type PersonalCharacteristicsObservationResolver struct {
	C *model.PersonalCharacteristicsObservation
}

//Id ..
func (r *PersonalCharacteristicsObservationResolver) Id() string {
	return r.C.Id
}

//Type ..
func (r *PersonalCharacteristicsObservationResolver) Type() model.PersonalCharacteristics {
	return r.C.Type
}

//Value ..
func (r *PersonalCharacteristicsObservationResolver) Value() string {
	return r.C.Value
}

//Code ..
func (r *PersonalCharacteristicsObservationResolver) Code() *CodableConceptResolver {
	return &CodableConceptResolver{r.C.Code}
}
