package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

/*==============================
Range Resolver
================================*/

// RangeResolver ..
type RangeResolver struct {
	L *model.Range
}

// Id ..
func (r *RangeResolver) Id() string {
	return r.L.Id
}

// Min ..
func (r *RangeResolver) Min() int32 {
	return r.L.Min
}

// Max ..
func (r *RangeResolver) Max() int32 {
	return r.L.Max
}
