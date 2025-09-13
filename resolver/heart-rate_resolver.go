package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
)

type HeartRateResolver struct {
	M *model.HeartRate
}

func (r *HeartRateResolver) Id() string {
	return r.M.Id
}
func (r *HeartRateResolver) Value() int32 {
	return r.M.Value
}

func (r *HeartRateResolver) Unit() string {
	return r.M.Unit
}
