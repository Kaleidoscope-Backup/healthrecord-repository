package service

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"

	"github.com/karte/mongo-lib/models"
	"github.com/karte/mongo-lib/mserver"
	"github.com/karte/mongo-lib/search"
	"gopkg.in/mgo.v2/bson"
)

// FindRecordsWithOrParams ...
func FindRecordsWithOrParams(params *map[string][]string, recordType interface{}, dal mserver.DataAccessLayer) (recordArr []interface{}, err error) {
	var bundle *models.Bundle

	//create a generic query object used by Mongo with ConsumerID as the query param
	q := createQueryObjectByOrParams(params, recordType)

	//run the actual search against Mongo
	bundle, err = dal.Search(url.URL{}, q, recordType) // the baseURL argument here does not matter
	if err != nil {
		return nil, errors.New(err.Error())
	}

	//Make sense of the returned elements from the search and make sure they are formatted to match the incoming recordType interface
	for _, entry := range bundle.Entry {
		record := reflect.New(reflect.TypeOf(recordType)).Interface()

		//unmarshal to interface
		bsonBytes, _ := bson.Marshal(entry.Resource)
		bson.Unmarshal(bsonBytes, &record)
		recordArr = append(recordArr, record)
	}

	return recordArr, nil
}

func createQueryObjectByOrParams(params *map[string][]string, recordType interface{}) search.Query {

	length := len(*params)
	query := search.Query{}

	if length > 0 {
		query.Resource = reflect.TypeOf(recordType).Elem().Name()

		var num int
		num = 0
		for k, v := range *params {
			fmt.Printf("%s -> %s\n", k, v)
			var arg string
			var i int
			i = 0
			for _, p := range v {
				arg += p
				i++
				if i < len(v) {
					arg += ","
				}
			}
			query.Query += k + "=" + arg
			num++
			if num < length {
				query.Query += "&"
			}
		}
	}

	return query
}

// FindRecords ...
func FindRecords(params *map[string]string, recordType interface{}, dal mserver.DataAccessLayer) (recordArr []interface{}, err error) {
	var bundle *models.Bundle

	//create a generic query object used by Mongo with ConsumerID as the query param
	q := createQueryObject(params, recordType)

	//run the actual search against Mongo
	bundle, err = dal.Search(url.URL{}, q, recordType) // the baseURL argument here does not matter
	if err != nil {
		return nil, errors.New(err.Error())
	}

	//Make sense of the returned elements from the search and make sure they are formatted to match the incoming recordType interface
	for _, entry := range bundle.Entry {
		record := reflect.New(reflect.TypeOf(recordType)).Interface()

		//unmarshal to interface
		bsonBytes, _ := bson.Marshal(entry.Resource)
		bson.Unmarshal(bsonBytes, &record)
		recordArr = append(recordArr, record)
	}

	return recordArr, nil
}

func createQueryObject(params *map[string]string, recordType interface{}) search.Query {

	length := len(*params)
	query := search.Query{}

	if length > 0 {
		query.Resource = reflect.TypeOf(recordType).Elem().Name()

		var num int
		num = 0
		for k, v := range *params {
			fmt.Printf("%s -> %s\n", k, v)
			query.Query += k + "=" + v
			num++
			if num < length {
				query.Query += "&"
			}
		}
	}

	return query
}
