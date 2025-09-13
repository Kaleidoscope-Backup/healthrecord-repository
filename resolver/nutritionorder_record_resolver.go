package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
NutritionOrderRecord Resolver
================================*/

// NutritionOrderRecordResolver ..
type NutritionOrderRecordResolver struct {
	HealthRecordResolver
	U *model.NutritionOrderRecord
}

// Id ..
func (r *NutritionOrderRecordResolver) Id() string {
	return r.U.Id
}

// Status ..
func (r *NutritionOrderRecordResolver) Status() model.NutritionOrderStatus {
	return r.U.Status
}

// Orderer ..
func (r *NutritionOrderRecordResolver) Orderer() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.U.Orderer}
}

// Product ..
func (r *NutritionOrderRecordResolver) Product() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{&r.U.Product}
}

// AllergyIntolerence array ..
func (r *NutritionOrderRecordResolver) AllergyIntolerence() *[]*ReferenceHealthRecordResolver {

	if r.U.AllergyIntolerence != nil {
		var crs []*ReferenceHealthRecordResolver
		var cs []model.ReferenceHealthRecord
		cs = *r.U.AllergyIntolerence

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ReferenceHealthRecord
				c = cs[i]
				if cr := ResolveReferenceHealthRecordResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// FoodPreferenceModifier ..
func (r *NutritionOrderRecordResolver) FoodPreferenceModifier() *[]model.Diet {
	return r.U.FoodPreferenceModifier
}

// ExcludeFoodModifier ..
func (r *NutritionOrderRecordResolver) ExcludeFoodModifier() *[]string {
	return r.U.ExcludeFoodModifier
}

// ExcludeFoodModifierCode array ..
func (r *NutritionOrderRecordResolver) ExcludeFoodModifierCode() *[]*CodableConceptResolver {

	if r.U.ExcludeFoodModifierCode != nil {
		var crs []*CodableConceptResolver
		var cs []model.CodableConcept
		cs = *r.U.ExcludeFoodModifierCode

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.CodableConcept
				c = cs[i]
				if cr := ResolveCodableConceptResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

// RouteOfAdministration ..
func (r *NutritionOrderRecordResolver) RouteOfAdministration() model.AdministrationRoute {
	return r.U.RouteOfAdministration
}

// RouteOfAdministrationCode ..
func (r *NutritionOrderRecordResolver) RouteOfAdministrationCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.U.RouteOfAdministrationCode}
}

// MaxVolumeToDeliver ..
func (r *NutritionOrderRecordResolver) MaxVolumeToDeliver() *int32 {
	return r.U.MaxVolumeToDeliver
}

// AdministrationInstruction ..
func (r *NutritionOrderRecordResolver) AdministrationInstruction() *string {
	return r.U.AdministrationInstruction
}
