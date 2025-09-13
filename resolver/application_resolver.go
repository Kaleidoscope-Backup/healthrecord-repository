package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
Application Resolver
================================*/

//ApplicationResolver ..
type ApplicationResolver struct {
	A *model.Application
}

//Id ..
func (r *ApplicationResolver) Id() string {
	return r.A.Id
}

//Name ..
func (r *ApplicationResolver) Name() string {
	return r.A.Name
}

//SupportEmail ..
func (r *ApplicationResolver) SupportEmail() *string {
	return r.A.SupportEmail
}

//DefaultLanguage ..
func (r *ApplicationResolver) DefaultLanguage() *model.Language {
	return r.A.DefaultLanguage
}

//Logo ..
func (r *ApplicationResolver) Logo() *string {
	return r.A.Logo
}

//Description ..
func (r *ApplicationResolver) Description() *string {
	return r.A.Description
}

//CallbackURL ..
func (r *ApplicationResolver) CallbackURL() *string {
	return r.A.CallbackURL
}

//Type ..
func (r *ApplicationResolver) Type() *model.ApplicationType {
	return r.A.Type
}

//Owner ..
func (r *ApplicationResolver) Owner() *ReferenceActorResolver {
	return &ReferenceActorResolver{&r.A.Owner}
}

//Attributes array ..
func (r *ApplicationResolver) Attributes() *[]*MetaDataResolver {

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
