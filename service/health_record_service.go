package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
)

/*==========================================================================================
Health Record service
==========================================================================================*/

//HealthRecordService is for creating health record
type HealthRecordService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewHealthRecordService ...
func NewHealthRecordService(dal mserver.DataAccessLayer, log *logging.Logger) *HealthRecordService {
	return &HealthRecordService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ..
func (u *HealthRecordService) FindByID(id string) (*model.HealthRecord, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Medication Record (if any) from Mongo
	p, err := u.dal.Get(id, &model.HealthRecord{})
	if err != nil {
		return nil, err
	}

	var hr *model.HealthRecord
	bsonBytes, _ := bson.Marshal(p)
	bson.Unmarshal(bsonBytes, &hr)

	return hr, nil
}

//FindByConsumerID ..
func (u *HealthRecordService) FindByConsumerID(id string) (*[]*model.HealthRecord, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching health Record (if any) from Mongo
	hrArr, err := FindHealthRecordsByConsumerID(id, &model.HealthRecord{}, u.dal)
	if err != nil {
		return nil, err
	}

	var healthRecordArr []*model.HealthRecord
	for _, hr := range hrArr {
		var healthRecord *model.HealthRecord
		bsonBytes, _ := bson.Marshal(hr)
		bson.Unmarshal(bsonBytes, &healthRecord)
		healthRecordArr = append(healthRecordArr, healthRecord)
	}

	return &healthRecordArr, nil
}
