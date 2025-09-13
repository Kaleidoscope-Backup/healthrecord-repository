package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
PersonalCharacteristicsObservationRecord Resolver
================================*/

// PersonalCharacteristicsObservationRecordResolver ..
type PersonalCharacteristicsObservationRecordResolver struct {
	HealthRecordResolver
	C *model.PersonalCharacteristicsObservationRecord
}

// Id ..
func (r *PersonalCharacteristicsObservationRecordResolver) Id() string {
	return r.C.Id
}

// Observations array ..
func (r *PersonalCharacteristicsObservationRecordResolver) Observations() *[]*PersonalCharacteristicsObservationResolver {

	if r.C.Observations != nil {
		var crs []*PersonalCharacteristicsObservationResolver
		var cs []model.PersonalCharacteristicsObservation
		cs = *r.C.Observations

		if r.C.Observations != nil && len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.PersonalCharacteristicsObservation
				c = cs[i]
				if cr := resolvePersonalCharacteristicsObservation(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolvePersonalCharacteristicsObservation(c *model.PersonalCharacteristicsObservation) *PersonalCharacteristicsObservationResolver {
	return &PersonalCharacteristicsObservationResolver{c}
}
