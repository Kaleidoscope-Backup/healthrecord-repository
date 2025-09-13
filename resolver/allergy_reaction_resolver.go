package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
AllergyReaction Resolver
================================*/

//AllergyReactionResolver ..
type AllergyReactionResolver struct {
	m *model.AllergyReaction
}

//Id ..
func (r *AllergyReactionResolver) Id() string {
	return r.m.Id
}

//Substance ..
func (r *AllergyReactionResolver) Substance() string {
	return r.m.Substance
}

//SubstanceCode ..
func (r *AllergyReactionResolver) SubstanceCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.m.SubstanceCode}
}

//Manifestation ..
func (r *AllergyReactionResolver) Manifestation() string {
	return r.m.Manifestation
}

//ManifestationCode ..
func (r *AllergyReactionResolver) ManifestationCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.m.ManifestationCode}
}

//ExposureRoute ..
func (r *AllergyReactionResolver) ExposureRoute() string {
	return r.m.ExposureRoute
}

//ExposureRouteCode ..
func (r *AllergyReactionResolver) ExposureRouteCode() *CodableConceptResolver {
	return &CodableConceptResolver{r.m.ExposureRouteCode}
}

//Description ..
func (r *AllergyReactionResolver) Description() *string {
	return r.m.Description
}

//Severity ..
func (r *AllergyReactionResolver) Severity() *model.Severity {
	return r.m.Severity
}
