package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

// ContactPointResolver ..
type ContactPointResolver struct {
	m *model.ContactPoint
}

// Id ..
func (r *ContactPointResolver) Id() string {
	return r.m.Id
}

// System ..
func (r *ContactPointResolver) System() string {
	return r.m.System
}

// Value ..
func (r *ContactPointResolver) Value() string {
	return r.m.Value
}

// Rank ..
func (r *ContactPointResolver) Rank() *int32 {
	return r.m.Rank
}

// Start ..
func (r *ContactPointResolver) Start() *util.Time {
	return r.m.Start
}

// End ..
func (r *ContactPointResolver) End() *util.Time {
	return r.m.End
}

// ResolveContactPointResolver ...
func ResolveContactPointResolver(contact *model.ContactPoint) *ContactPointResolver {
	if contact != nil {
		return &ContactPointResolver{contact}
	}

	return nil
}
