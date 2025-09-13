package service

import (
	"errors"
	"fmt"

	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
	"gopkg.in/mgo.v2/bson"
)

/*==========================================================================================
NutritionOrderRecord  service
==========================================================================================*/

//NutritionOrderRecordService is for creating habit
type NutritionOrderRecordService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewNutritionOrderRecordService Data Access Layer
func NewNutritionOrderRecordService(dal mserver.DataAccessLayer, log *logging.Logger) *NutritionOrderRecordService {
	return &NutritionOrderRecordService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ..
func (u *NutritionOrderRecordService) FindByID(id string) (*model.NutritionOrderRecord, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching NutritionOrderRecord (if any) from Mongo
	nor, err := u.dal.Get(id, &model.NutritionOrderRecord{})
	if err != nil {
		return nil, err
	}

	var nutritionOrder *model.NutritionOrderRecord
	bsonBytes, _ := bson.Marshal(nor)
	bson.Unmarshal(bsonBytes, &nutritionOrder)

	return nutritionOrder, nil
}

//FindByConsumerID ..
func (u *NutritionOrderRecordService) FindByConsumerID(id string) (*[]*model.NutritionOrderRecord, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching NutritionOrderRecord (if any) from Mongo
	norArr, err := FindHealthRecordsByConsumerID(id, &model.NutritionOrderRecord{}, u.dal)
	if err != nil {
		return nil, err
	}

	var nutritionOrderRecordArr []*model.NutritionOrderRecord
	for _, nor := range norArr {
		var nutritionOrderRecord *model.NutritionOrderRecord
		bsonBytes, _ := bson.Marshal(nor)
		bson.Unmarshal(bsonBytes, &nutritionOrderRecord)
		nutritionOrderRecordArr = append(nutritionOrderRecordArr, nutritionOrderRecord)
	}

	return &nutritionOrderRecordArr, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateNutritionOrderRecord will create a new record in Mongo using the Data Access Layer ...
func (u *NutritionOrderRecordService) CreateNutritionOrderRecord(nutritionOrderRecord *model.NutritionOrderRecord) (*model.NutritionOrderRecord, error) {

	if &nutritionOrderRecord.Status == nil || &nutritionOrderRecord.Orderer == nil || &nutritionOrderRecord.Product == nil || &nutritionOrderRecord.RouteOfAdministration == nil {
		return nil, errors.New("Missing a required field aborting before saving to the DB")
	}

	hError := ValidateHealthRecord(&nutritionOrderRecord.HealthRecord)
	if hError != nil {
		fmt.Printf("DB Create Failed For NutritionOrderRecord Record Error : %v", hError)
		return nil, hError
	}

	id, err := u.dal.Post(nutritionOrderRecord)
	if err != nil {
		fmt.Printf("DB Create Failed For NutritionOrderRecord Record Error : %v", err)
		return nil, err
	}

	errHr := u.dal.PostWithID(id, &nutritionOrderRecord.HealthRecord)
	if errHr != nil {
		return nil, err
	}

	PostRecord(nutritionOrderRecord, nutritionOrderRecord.ConsumerID, id)
	return nutritionOrderRecord, nil
}
