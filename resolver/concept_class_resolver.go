package resolver

import "gitlab.com/karte/healthrecord-repository/model"

/*==============================
ConceptClass Resolver
================================*/

//ConceptClassResolver ..
type ConceptClassResolver struct {
	C *model.ConceptClass
}

//Id ..
func (r *ConceptClassResolver) Id() string {
	return r.C.Id
}

//ExternalID ..
func (r *ConceptClassResolver) ExternalID() string {
	return r.C.ExternalID
}

//Name ..
func (r *ConceptClassResolver) Name() string {
	return r.C.Name
}

//Description ..
func (r *ConceptClassResolver) Description() *TextResolver {
	return &TextResolver{&r.C.Description}
}
