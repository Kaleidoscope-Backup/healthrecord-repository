package resolver

import "gitlab.com/karte/healthrecord-repository/model"

/*==============================
CodableConceptResolver Resolver
================================*/

//CodableConceptResolver ..
type CodableConceptResolver struct {
	C *model.CodableConcept
}

//Id ..
func (r *CodableConceptResolver) Id() string {
	return r.C.Id
}

//Text ..
func (r *CodableConceptResolver) Text() *string {
	return r.C.Text
}

//ConceptClass ..
func (r *CodableConceptResolver) ConceptClass() string {
	return r.C.ConceptClass
}

//Coding array ..
func (r *CodableConceptResolver) Coding() *[]*CodeResolver {

	if r.C.Coding != nil {
		var crs []*CodeResolver
		var cs []model.Code
		cs = *r.C.Coding

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.Code
				c = cs[i]
				if cr := resolveCodeResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}

func resolveCodeResolver(code *model.Code) *CodeResolver {
	return &CodeResolver{code}
}
