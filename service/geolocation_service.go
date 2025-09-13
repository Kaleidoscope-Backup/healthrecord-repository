package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
)

/*==========================================================================================
GeoLocation service
==========================================================================================*/

//GeoLocationService ..
type GeoLocationService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewGeoLocationService creates a new GeoLocation service that has all calls to the database, queries and mutations via the Data Access Layer
func NewGeoLocationService(dal mserver.DataAccessLayer, log *logging.Logger) *GeoLocationService {
	return &GeoLocationService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ..
func (u *GeoLocationService) FindByID(id string) (*model.GeoLocation, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching GeoLocation (if any) from Mongo
	p, err := u.dal.Get(id, &model.GeoLocation{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.GeoLocation
	//Convert BSON (byte) to JSON Fields
	var geoLocation *model.GeoLocation
	bsonBytes, _ := bson.Marshal(p)
	bson.Unmarshal(bsonBytes, &geoLocation)

	return geoLocation, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateGeoLocation will create a new geo location in Mongo using the Data Access Layer
func (u *GeoLocationService) CreateGeoLocation(location *model.GeoLocation) (*model.GeoLocation, error) {

	if &location.Latitude == nil || &location.Longitude == nil {
		return nil, errors.New("Missing a required field: aborting before saving to the DB")
	}
	_, err := u.dal.Post(location)
	if err != nil {
		return nil, err
	}

	return location, nil
}
