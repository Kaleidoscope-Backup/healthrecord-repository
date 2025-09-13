package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

// StrengthResolver ..
type StrengthResolver struct {
	M *model.Strength
}

// Id ..
func (r *StrengthResolver) Id() string {
	return r.M.Id
}

// Number ..
func (r *StrengthResolver) Number() int32 {
	return r.M.Number
}

// Unit ..
func (r *StrengthResolver) Unit() string {
	return r.M.Unit
}
