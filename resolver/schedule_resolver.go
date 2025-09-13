package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
Schedule Resolver
================================*/

//ScheduleResolver ..
type ScheduleResolver struct {
	S *model.Schedule
}

//Id ..
func (r *ScheduleResolver) Id() string {
	return r.S.Id
}

//Active ..
func (r *ScheduleResolver) Active() bool {
	return r.S.Active
}

//PlanningHorizon ..
func (r *ScheduleResolver) PlanningHorizon() *PeriodResolver {
	return &PeriodResolver{&r.S.PlanningHorizon}
}

//ServiceCategory ..
func (r *ScheduleResolver) ServiceCategory() *[]string {
	return r.S.ServiceCategory
}

//ServiceCategoryCode array ..
func (r *ScheduleResolver) ServiceCategoryCode() *[]*ClinicalCodeResolver {

	if r.S.ServiceCategoryCode != nil {

		var cprs []*ClinicalCodeResolver
		var cps []model.ClinicalCode
		cps = *r.S.ServiceCategoryCode

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

//Speciality ..
func (r *ScheduleResolver) Speciality() *[]string {
	return r.S.Speciality
}

//SpecialityCode array ..
func (r *ScheduleResolver) SpecialityCode() *[]*ClinicalCodeResolver {

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

//Actor array ..
func (r *ScheduleResolver) Actor() *[]*ReferenceActorResolver {

	if r.S.Actor != nil {

		var cprs []*ReferenceActorResolver
		var cps []model.ReferenceActor
		cps = *r.S.Actor

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

//Comment ..
func (r *ScheduleResolver) Comment() *string {
	return r.S.Comment
}
