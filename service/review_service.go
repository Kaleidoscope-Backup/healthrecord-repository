package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
)

/*==========================================================================================
ReviewService
==========================================================================================*/

//ReviewService ..
type ReviewService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewReviewService ..
func NewReviewService(dal mserver.DataAccessLayer, log *logging.Logger) *ReviewService {
	return &ReviewService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ..
func (u *ReviewService) FindByID(id string) (*model.Review, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching review (if any) from Mongo
	rv, err := u.dal.Get(id, &model.Review{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Slot
	//Convert BSON (byte) to JSON Fields
	var review *model.Review
	bsonBytes, _ := bson.Marshal(rv)
	bson.Unmarshal(bsonBytes, &review)

	return review, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateReview will create a new Review in Mongo using the Data Access Layer
func (u *ReviewService) CreateReview(review *model.Review) (*model.Review, error) {
	//Validate required fields on Model element are being passed in
	if review.Comment == "" {
		return nil, errors.New("Missing a required field Comment: aborting before saving to the DB")
	}

	if &review.Context == nil {
		return nil, errors.New("Missing a required field Context: aborting before saving to the DB")
	}

	if &review.By == nil {
		return nil, errors.New("Missing a required field By: aborting before saving to the DB")
	}

	id, err := u.dal.Post(review)
	if err != nil {
		return nil, err
	}

	review.Id = id
	return review, nil
}
