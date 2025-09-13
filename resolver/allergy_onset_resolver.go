package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
AllergyOnset Resolver
================================*/

// AllergyOnsetResolver ..
type AllergyOnsetResolver struct {
	m *model.AllergyOnset
}

// Id ..
func (r *AllergyOnsetResolver) Id() string {
	return r.m.Id
}

// OnsetDate ..
func (r *AllergyOnsetResolver) OnsetDate() *util.Time {
	return r.m.OnsetDate
}

// OnsetAge ..
func (r *AllergyOnsetResolver) OnsetAge() *string {
	return r.m.OnsetAge
}

// OnsetNote ..
func (r *AllergyOnsetResolver) OnsetNote() *string {
	return r.m.OnsetNote
}
