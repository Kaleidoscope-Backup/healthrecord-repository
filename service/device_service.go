package service

import (
	"errors"

	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
	"gopkg.in/mgo.v2/bson"
)

/*==========================================================================================
Device service
==========================================================================================*/

//DeviceService is for creating device
type DeviceService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewDeviceService creates a new device
func NewDeviceService(dal mserver.DataAccessLayer, log *logging.Logger) *DeviceService {
	return &DeviceService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ..
func (u *DeviceService) FindByID(id string) (*model.Device, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching consent (if any) from Mongo
	d, err := u.dal.Get(id, &model.Device{})
	if err != nil {
		return nil, err
	}

	var device *model.Device
	bsonBytes, _ := bson.Marshal(d)
	bson.Unmarshal(bsonBytes, &device)

	return device, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateDevice ...
func (u *DeviceService) CreateDevice(device *model.Device) (*model.Device, error) {

	_, err := u.dal.Post(device)
	if err != nil {
		return nil, err
	}

	return device, nil
}
