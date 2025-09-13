package resolver

import (
	"github.com/karte/healthrecord-repository/model"
)

// DosageResolver ..
type DosageResolver struct {
	M *model.Dosage
}

// Id ..
func (r *DosageResolver) Id() string {
	return r.M.Id
}

// Value ..
func (r *DosageResolver) Value() int32 {
	return r.M.Value
}

// Frequency ..
func (r *DosageResolver) Frequency() string {
	return r.M.Frequency
}

// Unit ..
func (r *DosageResolver) Unit() string {
	return r.M.Unit
}
