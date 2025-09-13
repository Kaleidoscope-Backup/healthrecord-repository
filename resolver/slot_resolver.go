package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
Slot Resolver
================================*/

// SlotResolver ..
type SlotResolver struct {
	S *model.Slot
}

// Id ..
func (r *SlotResolver) Id() string {
	return r.S.Id
}

// Status ..
func (r *SlotResolver) Status() model.SlotStatus {
	return r.S.Status
}

// OverBooked ..
func (r *SlotResolver) OverBooked() *bool {
	return r.S.OverBooked
}

// StatusCode ..
func (r *SlotResolver) StatusCode() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.S.StatusCode}
}

// Speciality ..
func (r *SlotResolver) Speciality() *[]string {
	return r.S.Speciality
}

// SpecialityCode array ..
func (r *SlotResolver) SpecialityCode() *[]*ClinicalCodeResolver {

	if r.S.SpecialityCode != nil {

		var cprs []*ClinicalCodeResolver
		var cps []model.ClinicalCode
		cps = *r.S.SpecialityCode

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.ClinicalCode
				cp = cps[i]
				if cpr := ResolveClinicalCodeResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

// ServiceType ..
func (r *SlotResolver) ServiceType() *[]string {
	return r.S.ServiceType
}

// ServiceTypeCode array ..
func (r *SlotResolver) ServiceTypeCode() *[]*ClinicalCodeResolver {

	if r.S.SpecialityCode != nil {

		var cprs []*ClinicalCodeResolver
		var cps []model.ClinicalCode
		cps = *r.S.ServiceTypeCode

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.ClinicalCode
				cp = cps[i]
				if cpr := ResolveClinicalCodeResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

// Schedule ..
func (r *SlotResolver) Schedule() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{&r.S.Schedule}
}

// Comment ..
func (r *SlotResolver) Comment() *string {
	return r.S.Comment
}

// Period ..
func (r *SlotResolver) Period() *PeriodResolver {
	return &PeriodResolver{r.S.Period}
}
