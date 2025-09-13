package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

// PractitionerResolver ..
type PractitionerResolver struct {
	ActorResolver
	M *model.Practitioner
}

// Id ..
func (r *PractitionerResolver) Id() string {
	return r.M.Id
}

// Speciality ..
func (r *PractitionerResolver) Speciality() string {
	return r.M.Speciality
}

// Qualification ..
func (r *PractitionerResolver) Qualification() string {
	return r.M.Qualification
}

// Organization ..
func (r *PractitionerResolver) Organization() string {
	return r.M.Organization
}

// Photo ..
func (r *PractitionerResolver) Photo() *string {
	return r.M.Photo
}

// Contacts ..
func (r *PractitionerResolver) Contacts() *[]*ContactPointResolver {
	if r.M.Contacts != nil {
		var crs []*ContactPointResolver
		var cs []model.ContactPoint
		cs = *r.M.Contacts

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.ContactPoint
				c = cs[i]
				if cr := ResolveContactPointResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
