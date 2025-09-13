package resolver

import (
	"gitlab.com/karte/healthrecord-repository/model"
)

/*==============================
Reference Resolver
================================*/

//ReferenceHealthRecordResolver ..
type ReferenceHealthRecordResolver struct {
	L *model.ReferenceHealthRecord
}

//Id ..
func (r *ReferenceHealthRecordResolver) Id() string {
	return r.L.Id
}

//Type ..
func (r *ReferenceHealthRecordResolver) Type() model.HealthRecordType {
	return r.L.Type
}

//ReferencedID ..
func (r *ReferenceHealthRecordResolver) ReferencedID() string {
	return r.L.ReferencedID
}
