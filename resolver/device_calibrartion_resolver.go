package resolver

import (
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
)

// DeviceCalibrationResolver ..
type DeviceCalibrationResolver struct {
	D *model.DeviceCalibration
}

// Id ..
func (r *DeviceCalibrationResolver) Id() string {
	return r.D.Id
}

// State ..
func (r *DeviceCalibrationResolver) State() model.DeviceMetricCalibrationState {
	return r.D.State
}

// Time ..
func (r *DeviceCalibrationResolver) Time() util.Time {
	return r.D.Time
}
