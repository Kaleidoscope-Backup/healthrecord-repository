package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

/*==============================
Period Resolver
================================*/

// PeriodResolver ..
type PeriodResolver struct {
	P *model.Period
}

// Id ..
func (r *PeriodResolver) Id() string {
	return r.P.Id
}

// Start ..
func (r *PeriodResolver) Start() util.Time {
	return r.P.Start
}

// End ..
func (r *PeriodResolver) End() util.Time {
	return r.P.End
}
