package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

//ContactResolver ..
type ContactResolver struct {
	m *model.Contact
}

//Id ..
func (r *ContactResolver) Id() string {
	return r.m.Id
}

//Name ..
func (r *ContactResolver) Name() string {
	return r.m.Name
}

//Relationship ..
func (r *ContactResolver) Relationship() string {
	return r.m.Relationship
}

//Phone ..
func (r *ContactResolver) Phone() string {
	return r.m.Phone
}

//Email ..
func (r *ContactResolver) Email() string {
	return r.m.Email
}
