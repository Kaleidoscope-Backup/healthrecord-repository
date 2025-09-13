package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

//ActorResolver ..
type ActorResolver struct {
	M *model.Actor
}

//FirstName ..
func (r *ActorResolver) FirstName() string {
	return r.M.FirstName
}

//LastName ..
func (r *ActorResolver) LastName() string {
	return r.M.LastName
}

//Email ..
func (r *ActorResolver) Email() string {
	return r.M.Email
}

//LanguagePreference ..
func (r *ActorResolver) LanguagePreference() *string {
	return r.M.LanguagePreference
}
