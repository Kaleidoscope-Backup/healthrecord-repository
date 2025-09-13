package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
ReferenceEntityResolver
================================*/

// ReferenceEntityResolver ..
type ReferenceEntityResolver struct {
	M *model.ReferenceEntity
}

// Id ..
func (r *ReferenceEntityResolver) Id() string {
	return r.M.Id
}

// EntityType ..
func (r *ReferenceEntityResolver) EntityType() model.EntityType {
	return r.M.EntityType
}

// EntityID ..
func (r *ReferenceEntityResolver) EntityID() string {
	return r.M.EntityID
}
