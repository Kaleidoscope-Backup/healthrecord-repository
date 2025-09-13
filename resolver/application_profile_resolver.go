package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
ApplicationProfile Resolver
================================*/

//ApplicationProfileResolver ..
type ApplicationProfileResolver struct {
	A *model.ApplicationProfile
}

//Name ..
func (r *ApplicationProfileResolver) Name() string {
	return r.A.Name
}

//Value ..
func (r *ApplicationProfileResolver) Value() string {
	return r.A.Value
}

//Attributes array ..
func (r *ApplicationProfileResolver) Attributes() *[]*MetaDataResolver {

	if r.A.Attributes != nil {
		var crs []*MetaDataResolver
		var cs []model.MetaData
		cs = *r.A.Attributes

		if len(cs) > 0 {
			for i := 0; i < len(cs); i++ {
				var c model.MetaData
				c = cs[i]
				if cr := ResolveMetaDataResolver(&c); cr != nil {
					crs = append(crs, cr)
				}
			}

			return &crs
		}
	}

	return nil
}
