package service

import (
	"errors"
	"strconv"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/constant"
	"github.com/Kaleidoscope-Backup/healthrecord-repository/model"
	"github.com/Kaleidoscope-Backup/mongo-lib/mserver"
	"github.com/globalsign/mgo/bson"
	logging "github.com/op/go-logging"
)

/*==========================================================================================
CustomerFeedback service
==========================================================================================*/

// CustomerFeedbackService is for creating customer feedback
type CustomerFeedbackService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewCustomerFeedbackService ...
func NewCustomerFeedbackService(dal mserver.DataAccessLayer, log *logging.Logger) *CustomerFeedbackService {
	return &CustomerFeedbackService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *CustomerFeedbackService) FindByID(id string) (*model.CustomerFeedback, error) {

	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching customer feedback from Mongo
	c, err := u.dal.Get(id, &model.CustomerFeedback{})
	if err != nil {
		return nil, err
	}

	var feedback *model.CustomerFeedback
	bsonBytes, _ := bson.Marshal(c)
	bson.Unmarshal(bsonBytes, &feedback)

	return feedback, nil
}

// FindByApplicationID ...
func (u *CustomerFeedbackService) FindByApplicationID(applicationID *string) (*[]*model.CustomerFeedback, error) {
	if applicationID == nil {
		return nil, errors.New("Missing parameter - application ID should be provided")
	}

	var params map[string]string
	params = map[string]string{}

	params["applicationID"] = string(*applicationID)
	params["_count"] = strconv.Itoa(int(constant.MAX_RECORD_FETCH_COUNT))

	//find the matching relationship Record (if any) from Mongo
	fdArr, err := FindRecords(&params, &model.CustomerFeedback{}, u.dal)
	if err != nil {
		return nil, err
	}

	var customerFeedbackArr []*model.CustomerFeedback
	for _, fd := range fdArr {
		var feedback *model.CustomerFeedback
		bsonBytes, _ := bson.Marshal(fd)
		bson.Unmarshal(bsonBytes, &feedback)
		customerFeedbackArr = append(customerFeedbackArr, feedback)
	}

	return &customerFeedbackArr, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateCustomerFeedback ...
func (u *CustomerFeedbackService) CreateCustomerFeedback(feedback *model.CustomerFeedback) (*model.CustomerFeedback, error) {
	id, err := u.dal.Post(feedback)
	if err != nil {
		return nil, err
	}

	feedback.Id = id
	return feedback, nil
}
