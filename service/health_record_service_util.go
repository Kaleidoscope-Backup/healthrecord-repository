package service

import (
	"errors"
	"net/url"
	"reflect"

	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/models"
	"gitlab.com/karte/mongo-lib/mserver"
	"gitlab.com/karte/mongo-lib/search"
	"gopkg.in/mgo.v2/bson"
)

//ValidateHealthRecord ...
func ValidateHealthRecord(healthRecord *model.HealthRecord) error {

	if &healthRecord.TransactionType == nil {
		return errors.New("Missing a required field TransactionType aborting before saving to the DB")
	}

	if healthRecord.Name == "" {
		return errors.New("Missing a required field Name aborting before saving to the DB")
	}

	if healthRecord.ConsumerID == "" {
		return errors.New("Missing a required field Consumer ID aborting before saving to the DB")
	}

	if healthRecord.Occurred.IsZero() {
		return errors.New("Missing a required field Occurred aborting before saving to the DB")
	}

	if healthRecord.Created.IsZero() {
		return errors.New("Missing a required field Created aborting before saving to the DB")
	}

	if healthRecord.Source == "" {
		return errors.New("Missing a required field Source aborting before saving to the DB")
	}

	return nil
}

//FindHealthRecordsByConsumerID does a search across the mongo database to surface all records associated to that Consumer
func FindHealthRecordsByConsumerID(consumerID string, recordType interface{}, dal mserver.DataAccessLayer) (recordArr []interface{}, err error) {
	var bundle *models.Bundle

	//create a generic query object used by Mongo with ConsumerID as the query param
	q := createHealthRecordSearchObject(consumerID, recordType)

	//run the actual search against Mongo
	bundle, err = dal.Search(url.URL{}, q, recordType) // the baseURL argument here does not matter
	if err != nil {
		return nil, errors.New(err.Error())
	}

	//Make sense of the returned elements from the search and make sure they are formatted to match the incoming recordType interface
	for _, entry := range bundle.Entry {
		record := reflect.New(reflect.TypeOf(recordType)).Interface()
		//unmarshal to bundle entry component
		var r *models.BundleEntryComponent
		bsonBytes, _ := bson.Marshal(entry)
		bson.Unmarshal(bsonBytes, &r)
		//unmarshal to interface
		bsonBytes, _ = bson.Marshal(r.Resource)
		bson.Unmarshal(bsonBytes, &record)
		recordArr = append(recordArr, record)
	}

	return recordArr, nil
}

func createHealthRecordSearchObject(consumerID string, recordType interface{}) search.Query {
	query := search.Query{}
	query.Resource = reflect.TypeOf(recordType).Elem().Name()
	query.Query = "consumerID=" + consumerID
	query.Query += "&_sort:desc=_lastUpdated"
	return query
}
