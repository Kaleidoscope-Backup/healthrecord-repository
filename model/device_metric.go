package model

import (
	"github.com/karte/healthrecord-repository/util"
	"github.com/karte/mongo-lib/models"
)

// DeviceMetricOperationalStatus enum ...
type DeviceMetricOperationalStatus string

const (
	//DEVICE_METRIC_STATUS_ON ...
	DEVICE_METRIC_STATUS_ON DeviceMetricOperationalStatus = "DEVICE_METRIC_STATUS_ON"

	//DEVICE_METRIC_STATUS_OFF ...
	DEVICE_METRIC_STATUS_OFF DeviceMetricOperationalStatus = "DEVICE_METRIC_STATUS_OFF"

	//DEVICE_METRIC_STATUS_STANDBY ...
	DEVICE_METRIC_STATUS_STANDBY DeviceMetricOperationalStatus = "DEVICE_METRIC_STATUS_STANDBY"

	//DEVICE_METRIC_STATUS_ENTERED_IN_ERROR ...
	DEVICE_METRIC_STATUS_ENTERED_IN_ERROR DeviceMetricOperationalStatus = "DEVICE_METRIC_STATUS_ENTERED_IN_ERROR"
)

// DeviceMetricCategory enum ...
type DeviceMetricCategory string

const (
	//DEVICE_METRIC_MEASUREMENT ...
	DEVICE_METRIC_MEASUREMENT DeviceMetricCategory = "DEVICE_METRIC_MEASUREMENT"

	//DEVICE_METRIC_SETTING ...
	DEVICE_METRIC_SETTING DeviceMetricCategory = "DEVICE_METRIC_SETTING"

	//DEVICE_METRIC_CALCULATION ...
	DEVICE_METRIC_CALCULATION DeviceMetricCategory = "DEVICE_METRIC_CALCULATION"

	//DEVICE_METRIC_UNSPECIFIED ...
	DEVICE_METRIC_UNSPECIFIED DeviceMetricCategory = "DEVICE_METRIC_UNSPECIFIED"
)

// DeviceMetricCalibrationState enum ...
type DeviceMetricCalibrationState string

const (
	//DEVICE_METRIC_CALIBRATION_REQUIRED ...
	DEVICE_METRIC_CALIBRATION_REQUIRED DeviceMetricCalibrationState = "DEVICE_METRIC_CALIBRATION_REQUIRED"

	//DEVICE_METRIC_CALIBRATED ...
	DEVICE_METRIC_CALIBRATED DeviceMetricCalibrationState = "DEVICE_METRIC_CALIBRATED"

	//DEVICE_METRIC_CALIBRATION_NOT_REQUIRED ...
	DEVICE_METRIC_CALIBRATION_NOT_REQUIRED DeviceMetricCalibrationState = "DEVICE_METRIC_CALIBRATION_NOT_REQUIRED"

	//DEVICE_METRIC_CALIBRATION_UNSPECIFIED ...
	DEVICE_METRIC_CALIBRATION_UNSPECIFIED DeviceMetricCalibrationState = "DEVICE_METRIC_CALIBRATION_UNSPECIFIED"
)

// DeviceCalibrationInput ...
type DeviceCalibrationInput struct {
	State DeviceMetricCalibrationState `json:"state"`
	Time  util.Time                    `json:"time"`
}

// DeviceCalibration ...
type DeviceCalibration struct {
	Id    string                       `json:"id" bson:"_id"`
	State DeviceMetricCalibrationState `json:"state" bson:"state"`
	Time  util.Time                    `json:"time" bson:"time"`
}

// DeviceMetricCreate ...
type DeviceMetricCreate struct {
	Type              string                         `json:"type"`
	Unit              string                         `json:"unit"`
	Source            ReferenceEntityInput           `json:"source"`
	OperationalStatus *DeviceMetricOperationalStatus `json:"operationalStatus"`
	Category          *DeviceMetricCategory          `json:"category"`
	Calibration       *DeviceCalibrationInput        `json:"calibration"`
	MeasurementPeriod *TimingInput                   `json:"measurementPeriod"`
}

// DeviceMetric ...
type DeviceMetric struct {
	Id                string                         `json:"id" bson:"_id"`
	Type              string                         `json:"type" bson:"type"`
	Unit              string                         `json:"unit" bson:"unit"`
	Source            ReferenceEntity                `json:"source" bson:"source"`
	OperationalStatus *DeviceMetricOperationalStatus `json:"operationalStatus" bson:"operationalStatus"`
	Category          *DeviceMetricCategory          `json:"category" bson:"category"`
	Calibration       *DeviceCalibration             `json:"calibration" bson:"calibration"`
	MeasurementPeriod *Timing                        `json:"measurementPeriod" bson:"measurementPeriod"`
	Meta              *models.Meta                   //MUST INCLUDE to capture meta data (including timestamps) and is used in Mongo Lib for reflection
}
