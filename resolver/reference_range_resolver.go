package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
ReferenceRange Resolver
================================*/

//ReferenceRangeResolver ..
type ReferenceRangeResolver struct {
	L *model.ReferenceRange
}

//Id ..
func (r *ReferenceRangeResolver) Id() string {
	return r.L.Id
}

//LowerLimit ..
func (r *ReferenceRangeResolver) LowerLimit() *int32 {
	return r.L.LowerLimit
}

//HigherLimit ..
func (r *ReferenceRangeResolver) HigherLimit() *int32 {
	return r.L.HigherLimit
}

//RangeUnit ..
func (r *ReferenceRangeResolver) RangeUnit() *string {
	return r.L.RangeUnit
}

//Range ..
func (r *ReferenceRangeResolver) Range() *RangeResolver {
	return &RangeResolver{r.L.Range}
}

//AgeMin ..
func (r *ReferenceRangeResolver) AgeMin() *int32 {
	return r.L.AgeMin
}

//AgeMax ..
func (r *ReferenceRangeResolver) AgeMax() *int32 {
	return r.L.AgeMax
}

//AgeUnit ..
func (r *ReferenceRangeResolver) AgeUnit() *string {
	return r.L.AgeUnit
}

//AgeGroup ..
func (r *ReferenceRangeResolver) AgeGroup() *[]model.AgeGroup {
	return r.L.AgeGroup
}

//Type ..
func (r *ReferenceRangeResolver) Type() model.ReferenceRangeType {
	return r.L.Type
}

//AppliesTo ..
func (r *ReferenceRangeResolver) AppliesTo() *[]string {
	return r.L.AppliesTo
}

//AppliesToCode ..
func (r *ReferenceRangeResolver) AppliesToCode() *[]*ClinicalCodeResolver {

	if r.L.AppliesToCode != nil {

		var cprs []*ClinicalCodeResolver
		var cps []model.ClinicalCode
		cps = *r.L.AppliesToCode

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
