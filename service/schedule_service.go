package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/mongo-lib/mserver"
)

/*==========================================================================================
ScheduleService
==========================================================================================*/

//ScheduleService ..
type ScheduleService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewScheduleService ..
func NewScheduleService(dal mserver.DataAccessLayer, log *logging.Logger) *ScheduleService {
	return &ScheduleService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ..
func (u *ScheduleService) FindByID(id string) (*model.Schedule, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching schedule (if any) from Mongo
	prd, err := u.dal.Get(id, &model.Schedule{})
	if err != nil {
		return nil, err
	}

	//we need to unmarshal the result from type bson.M{} to model.Dosage
	//Convert BSON (byte) to JSON Fields
	var schedule *model.Schedule
	bsonBytes, _ := bson.Marshal(prd)
	bson.Unmarshal(bsonBytes, &schedule)

	return schedule, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateSchedule will create a new Schedule in Mongo using the Data Access Layer
func (u *ScheduleService) CreateSchedule(schedule *model.Schedule) (*model.Schedule, error) {
	//Validate required fields on Model element are being passed in
	if &schedule.Active == nil {
		return nil, errors.New("Missing a required field Status: aborting before saving to the DB")
	}

	if &schedule.PlanningHorizon == nil {
		return nil, errors.New("Missing a required field: PlanningHorizon aborting before saving to the DB")
	}

	id, err := u.dal.Post(schedule)
	if err != nil {
		return nil, err
	}

	schedule.Id = id
	return schedule, nil
}
