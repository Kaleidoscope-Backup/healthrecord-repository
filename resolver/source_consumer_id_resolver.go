package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
)

/*==============================
SourceConsumerIDResolver
================================*/

// SourceConsumerIDResolver ..
type SourceConsumerIDResolver struct {
	M *model.SourceConsumerID
}

// Id ..
func (r *SourceConsumerIDResolver) Id() string {
	return r.M.Id
}

// System ..
func (r *SourceConsumerIDResolver) System() string {
	return r.M.System
}

// Value ..
func (r *SourceConsumerIDResolver) Value() string {
	return r.M.Value
}

// Assigner ..
func (r *SourceConsumerIDResolver) Assigner() string {
	return r.M.Assigner
}

// Use ..
func (r *SourceConsumerIDResolver) Use() *model.SourceConsumerIDUse {
	return r.M.Use
}

// Type ..
func (r *SourceConsumerIDResolver) Type() *model.SourceConsumerIDType {
	return r.M.Type
}

// Start ..
func (r *SourceConsumerIDResolver) Start() *util.Time {
	return r.M.Start
}

// End ..
func (r *SourceConsumerIDResolver) End() *util.Time {
	return r.M.End
}
