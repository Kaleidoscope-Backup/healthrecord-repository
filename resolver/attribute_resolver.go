package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
Attribute Resolver
================================*/

//AttributeResolver ..
type AttributeResolver struct {
	A *model.Attribute
}

//Id ..
func (r *AttributeResolver) Id() string {
	return r.A.Id
}

//Name ..
func (r *AttributeResolver) Name() string {
	return r.A.Name
}

//Value ..
func (r *AttributeResolver) Value() *ValueResolver {
	return &ValueResolver{&r.A.Value}
}
