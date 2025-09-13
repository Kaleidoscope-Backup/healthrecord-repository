package resolver

import (
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
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
