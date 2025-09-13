package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

/*==============================
Reference Actor Resolver
================================*/

// ReferenceActorResolver ..
type ReferenceActorResolver struct {
	L *model.ReferenceActor
}

// Id ..
func (r *ReferenceActorResolver) Id() string {
	return r.L.Id
}

// ActorType ..
func (r *ReferenceActorResolver) ActorType() model.ActorType {
	return r.L.ActorType
}

// ActorID ..
func (r *ReferenceActorResolver) ActorID() string {
	return r.L.ActorID
}
