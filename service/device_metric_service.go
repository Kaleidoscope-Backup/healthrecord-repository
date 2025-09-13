package service

import (
	"errors"

	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/mongo-lib/mserver"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2/bson"
)

/*==========================================================================================
Device Metric service
==========================================================================================*/

// DeviceMetricService is for creating device
type DeviceMetricService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewDeviceMetricService creates a new device metric
func NewDeviceMetricService(dal mserver.DataAccessLayer, log *logging.Logger) *DeviceMetricService {
	return &DeviceMetricService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *DeviceMetricService) FindByID(id string) (*model.DeviceMetric, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching consent (if any) from Mongo
	d, err := u.dal.Get(id, &model.DeviceMetric{})
	if err != nil {
		return nil, err
	}

	var deviceMetric *model.DeviceMetric
	bsonBytes, _ := bson.Marshal(d)
	bson.Unmarshal(bsonBytes, &deviceMetric)

	return deviceMetric, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateDeviceMetric ...
func (u *DeviceMetricService) CreateDeviceMetric(deviceMetric *model.DeviceMetric) (*model.DeviceMetric, error) {

	_, err := u.dal.Post(deviceMetric)
	if err != nil {
		return nil, err
	}

	return deviceMetric, nil
}
