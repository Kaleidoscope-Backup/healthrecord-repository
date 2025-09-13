package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
PersonalCharacteristicsObservationRecord Resolver
================================*/

// VitalObservationRecordResolver ..
type VitalObservationRecordResolver struct {
	HealthRecordResolver
	C *model.VitalObservationRecord
}

// Id ..
func (r *VitalObservationRecordResolver) Id() string {
	return r.C.Id
}

// Observations array ..
func (r *VitalObservationRecordResolver) Observations() *[]*VitalResolver {

	if r.C.Observations != nil {
		var crs []*VitalResolver
		var cs []model.Vital
		cs = *r.C.Observations

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Vital
				c = cs[i]
				if cr := resolveVitalResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolveVitalResolver(c *model.Vital) *VitalResolver {
	return &VitalResolver{c}
}
