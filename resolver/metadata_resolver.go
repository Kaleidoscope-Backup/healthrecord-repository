package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
MetaData Resolver
================================*/

//MetaDataResolver ..
type MetaDataResolver struct {
	A *model.MetaData
}

//Id ..
func (r *MetaDataResolver) Id() string {
	return r.A.Id
}

//Name ..
func (r *MetaDataResolver) Name() string {
	return r.A.Name
}

//Value ..
func (r *MetaDataResolver) Value() string {
	return r.A.Value
}

//Attributes ..
func (r *MetaDataResolver) Attributes() *[]*MetaDataResolver {

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
