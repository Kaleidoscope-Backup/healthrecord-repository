package service

import (
	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
)

type HeartRateService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

func NewHeartRateService(dal mserver.DataAccessLayer, log *logging.Logger) *HeartRateService {
	return &HeartRateService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

func (u *HeartRateService) FindByID(id string) (*model.HeartRate, error) {
	//find the matching Heart Rate (if any) from Mongo
	hr, err := u.dal.Get(id, &model.HeartRate{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.HeartRate
	//Convert BSON (byte) to JSON Fields
	var heartRate *model.HeartRate
	bsonBytes, _ := bson.Marshal(hr)
	bson.Unmarshal(bsonBytes, &heartRate)

	return heartRate, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

func (u *HeartRateService) CreateHeartRate(heartRate *model.HeartRate) (*model.HeartRate, error) {
	_, err := u.dal.Post(heartRate)
	if err != nil {
		return nil, err
	}

	return heartRate, nil
}
