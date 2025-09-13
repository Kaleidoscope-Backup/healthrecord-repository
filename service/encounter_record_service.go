package service

import (
	"errors"
	"fmt"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/mongo-lib/mserver"
	"github.com/globalsign/mgo/bson"
	logging "github.com/op/go-logging"
)

/*==========================================================================================
EncounterRecordService
==========================================================================================*/

// EncounterRecordService ..
type EncounterRecordService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewEncounterRecordService ..
func NewEncounterRecordService(dal mserver.DataAccessLayer, log *logging.Logger) *EncounterRecordService {
	return &EncounterRecordService{dal: dal, log: log}
}

// FindById ...
func (ers *EncounterRecordService) FindById(id string) (*model.EncounterRecord, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching EncounterRecord (if any) from Mongo
	cct, err := ers.dal.Get(id, &model.EncounterRecord{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.EncounterRecord
	//Convert BSON (byte) to JSON Fields
	var encounterRecord *model.EncounterRecord
	bsonBytes, _ := bson.Marshal(cct)
	bson.Unmarshal(bsonBytes, &encounterRecord)

	return encounterRecord, nil
}

// FindByConsumerID ..
func (ers *EncounterRecordService) FindByConsumerID(id string) (*[]*model.EncounterRecord, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Medication Record (if any) from Mongo
	erArr, err := FindHealthRecordsByConsumerID(id, &model.EncounterRecord{}, ers.dal)
	if err != nil {
		return nil, err
	}

	var encounterRecordArr []*model.EncounterRecord
	for _, er := range erArr {
		var encounterRecord *model.EncounterRecord
		bsonBytes, _ := bson.Marshal(er)
		bson.Unmarshal(bsonBytes, &encounterRecord)
		encounterRecordArr = append(encounterRecordArr, encounterRecord)
	}

	return &encounterRecordArr, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateEncounterRecord will create a new EncounterRecord in Mongo using the Data Access Layer
func (ers *EncounterRecordService) CreateEncounterRecord(encounterRecord *model.EncounterRecord) (*model.EncounterRecord, error) {

	hError := ValidateHealthRecord(&encounterRecord.HealthRecord)
	if hError != nil {
		fmt.Printf("DB Create Failed For Allergy Record Error : %v", hError)
		return nil, hError
	}

	id, err := ers.dal.Post(encounterRecord)
	if err != nil {
		return nil, err
	}

	encounterRecord.Id = id
	errHr := ers.dal.PostWithID(id, &encounterRecord.HealthRecord)
	if errHr != nil {
		return nil, err
	}

	PostRecord(encounterRecord, encounterRecord.ConsumerID, id)
	return encounterRecord, nil
}
