package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
Abatement Resolver
================================*/

// AbatementResolver ..
type AbatementResolver struct {
	M *model.Abatement
}

// Id ..
func (r *AbatementResolver) Id() string {
	return r.M.Id
}

// Abatement ..
func (r *AbatementResolver) Abatement() *bool {
	return r.M.Abatement
}

// Date ..
func (r *AbatementResolver) Date() *util.Time {
	return r.M.Date
}

// Age ..
func (r *AbatementResolver) Age() *string {
	return r.M.Age
}

// Note ..
func (r *AbatementResolver) Note() *string {
	return r.M.Note
}
