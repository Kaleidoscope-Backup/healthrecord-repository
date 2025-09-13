package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
Relationship Resolver
================================*/

//RelationshipResolver ...
type RelationshipResolver struct {
	R *model.Relationship
}

//Id ...
func (r *RelationshipResolver) Id() string {
	return r.R.Id
}

//Active ...
func (r *RelationshipResolver) Active() bool {
	return r.R.Active
}

//From ...
func (r *RelationshipResolver) From() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.R.From}
}

//To ...
func (r *RelationshipResolver) To() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.R.To}
}

//Label ...
func (r *RelationshipResolver) Label() string {
	return r.R.Label
}

//Type ...
func (r *RelationshipResolver) Type() *model.RelationshipType {
	return r.R.Type
}

//Code ...
func (r *RelationshipResolver) Code() *ClinicalCodeResolver {
	return &ClinicalCodeResolver{r.R.Code}
}

//Consent ...
func (r *RelationshipResolver) Consent() *string {
	return r.R.Consent
}

//Period ...
func (r *RelationshipResolver) Period() *PeriodResolver {
	return &PeriodResolver{r.R.Period}
}

//AdditionalData array ..
func (r *RelationshipResolver) AdditionalData() *[]*AttributeResolver {

	if r.R.AdditionalData != nil {
		var crs []*AttributeResolver
		var cs []model.Attribute
		cs = *r.R.AdditionalData

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Attribute
				c = cs[i]
				if cr := ResolveAttributeResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
