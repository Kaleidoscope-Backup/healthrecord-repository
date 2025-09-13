package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
AppointmentRecord Resolver
================================*/

// AppointmentRecordResolver ..
type AppointmentRecordResolver struct {
	HealthRecordResolver
	A *model.AppointmentRecord
}

// Id ..
func (r *AppointmentRecordResolver) Id() string {
	return r.A.Id
}

// Status ..
func (r *AppointmentRecordResolver) Status() model.AppointmentStatus {
	return r.A.Status
}

// StatusCode ..
func (r *AppointmentRecordResolver) StatusCode() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.A.StatusCode}
}

// Speciality ..
func (r *AppointmentRecordResolver) Speciality() *[]string {
	return r.A.Speciality
}

// SpecialityCode array ..
func (r *AppointmentRecordResolver) SpecialityCode() *[]*ClinicalCodeResolver {

	if r.A.SpecialityCode != nil {

		var cprs []*ClinicalCodeResolver
		var cps []model.ClinicalCode
		cps = *r.A.SpecialityCode

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

// AppointmentType ..
func (r *AppointmentRecordResolver) AppointmentType() model.AppointmentType {
	return r.A.AppointmentType
}

// ServiceCategory ..
func (r *AppointmentRecordResolver) ServiceCategory() *[]string {
	return r.A.ServiceCategory
}

// ServiceCategoryCode array ..
func (r *AppointmentRecordResolver) ServiceCategoryCode() *[]*ClinicalCodeResolver {

	if r.A.ServiceCategoryCode != nil {

		var cprs []*ClinicalCodeResolver
		var cps []model.ClinicalCode
		cps = *r.A.ServiceCategoryCode

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

// Reason ..
func (r *AppointmentRecordResolver) Reason() *[]string {
	return r.A.Reason
}

// ReasonCode array ..
func (r *AppointmentRecordResolver) ReasonCode() *[]*ClinicalCodeResolver {

	if r.A.ReasonCode != nil {

		var cprs []*ClinicalCodeResolver
		var cps []model.ClinicalCode
		cps = *r.A.ReasonCode

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

// Indication array ..
func (r *AppointmentRecordResolver) Indication() *[]*ReferenceHealthRecordResolver {

	if r.A.Indication != nil {

		var cprs []*ReferenceHealthRecordResolver
		var cps []model.ReferenceHealthRecord
		cps = *r.A.Indication

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.ReferenceHealthRecord
				cp = cps[i]
				if cpr := ResolveReferenceHealthRecordResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

// IncomingReferral array ..
func (r *AppointmentRecordResolver) IncomingReferral() *[]*ReferenceEntityResolver {

	if r.A.IncomingReferral != nil {

		var cprs []*ReferenceEntityResolver
		var cps []model.ReferenceEntity
		cps = *r.A.IncomingReferral

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.ReferenceEntity
				cp = cps[i]
				if cpr := ResolveReferenceEntityResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

// Priority ..
func (r *AppointmentRecordResolver) Priority() *model.Priority {
	return r.A.Priority
}

// MinutesDuration ..
func (r *AppointmentRecordResolver) MinutesDuration() *int32 {
	return r.A.MinutesDuration
}

// Slot array ..
func (r *AppointmentRecordResolver) Slot() *[]*ReferenceEntityResolver {

	if r.A.Slot != nil {

		var cprs []*ReferenceEntityResolver
		var cps []model.ReferenceEntity
		cps = *r.A.Slot

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.ReferenceEntity
				cp = cps[i]
				if cpr := ResolveReferenceEntityResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

// Comment ..
func (r *AppointmentRecordResolver) Comment() *string {
	return r.A.Comment
}

// Participants array ..
func (r *AppointmentRecordResolver) Participants() *[]*ReferenceActorResolver {

	if r.A.Participants != nil {

		var cprs []*ReferenceActorResolver
		var cps []model.ReferenceActor
		cps = *r.A.Participants

		if len(cps) > 0 {
			for i := 0; i < len(cps); i++ {
				var cp model.ReferenceActor
				cp = cps[i]
				if cpr := ResolveReferenceActorResolver(&cp); cpr != nil {
					cprs = append(cprs, cpr)
				}
			}

			return &cprs
		}
	}

	return nil
}

// RequestedPeriod ..
func (r *AppointmentRecordResolver) RequestedPeriod() *PeriodResolver {
	return &PeriodResolver{&r.A.RequestedPeriod}
}
