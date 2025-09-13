package service

import (
	"errors"

	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/mongo-lib/mserver"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2/bson"
)

/*==========================================================================================
Device Data Source service
==========================================================================================*/

// DeviceDataSourceService is for creating device data source
type DeviceDataSourceService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewDeviceDataSourceService creates a new device data source
func NewDeviceDataSourceService(dal mserver.DataAccessLayer, log *logging.Logger) *DeviceDataSourceService {
	return &DeviceDataSourceService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *DeviceDataSourceService) FindByID(id string) (*model.DeviceDataSource, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching consent (if any) from Mongo
	d, err := u.dal.Get(id, &model.DeviceDataSource{})
	if err != nil {
		return nil, err
	}

	var deviceDataSource *model.DeviceDataSource
	bsonBytes, _ := bson.Marshal(d)
	bson.Unmarshal(bsonBytes, &deviceDataSource)

	return deviceDataSource, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateDeviceDataSource ...
func (u *DeviceDataSourceService) CreateDeviceDataSource(deviceDataSource *model.DeviceDataSource) (*model.DeviceDataSource, error) {

	_, err := u.dal.Post(deviceDataSource)
	if err != nil {
		return nil, err
	}

	return deviceDataSource, nil
}
