package resolver

import "github.com/karte/healthrecord-repository/model"

// DeviceMetricResolver ..
type DeviceMetricResolver struct {
	D *model.DeviceMetric
}

// Id ..
func (r *DeviceMetricResolver) Id() string {
	return r.D.Id
}

// Type ..
func (r *DeviceMetricResolver) Type() string {
	return r.D.Type
}

// Unit ..
func (r *DeviceMetricResolver) Unit() string {
	return r.D.Unit
}

// Source ..
func (r *DeviceMetricResolver) Source() *ReferenceEntityResolver {
	return &ReferenceEntityResolver{&r.D.Source}
}

// OperationalStatus ..
func (r *DeviceMetricResolver) OperationalStatus() *model.DeviceMetricOperationalStatus {
	return r.D.OperationalStatus
}

// Category ..
func (r *DeviceMetricResolver) Category() *model.DeviceMetricCategory {
	return r.D.Category
}

// Calibration ..
func (r *DeviceMetricResolver) Calibration() *DeviceCalibrationResolver {
	return &DeviceCalibrationResolver{r.D.Calibration}
}

// MeasurementPeriod ..
func (r *DeviceMetricResolver) MeasurementPeriod() *TimingResolver {
	return &TimingResolver{r.D.MeasurementPeriod}
}
