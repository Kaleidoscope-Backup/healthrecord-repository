package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
Onset Resolver
================================*/

// OnsetResolver ..
type OnsetResolver struct {
	M *model.Onset
}

// Id ..
func (r *OnsetResolver) Id() string {
	return r.M.Id
}

// Date ..
func (r *OnsetResolver) Date() *util.Time {
	return r.M.Date
}

// Age ..
func (r *OnsetResolver) Age() *string {
	return r.M.Age
}

// Note ..
func (r *OnsetResolver) Note() *string {
	return r.M.Note
}
