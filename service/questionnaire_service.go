package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/mongo-lib/mserver"
	"github.com/op/go-logging"
)

/*==========================================================================================
Questionnaire service
==========================================================================================*/

// QuestionnaireService is for creating Questionnaire
type QuestionnaireService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewQuestionnaireService creates a new Questionnaire service
func NewQuestionnaireService(dal mserver.DataAccessLayer, log *logging.Logger) *QuestionnaireService {
	return &QuestionnaireService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *QuestionnaireService) FindByID(id string) (*model.Questionnaire, error) {

	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Contact (if any) from Mongo
	p, err := u.dal.Get(id, &model.Questionnaire{})
	if err != nil {
		return nil, err
	}

	var questionnaire *model.Questionnaire
	bsonBytes, _ := bson.Marshal(p)
	bson.Unmarshal(bsonBytes, &questionnaire)

	return questionnaire, nil
}

// FindQuestionnaires ..
func (u *QuestionnaireService) FindQuestionnaires(param *model.QuestionnaireQueryParam) (*[]*model.Questionnaire, error) {

	//find the matching Contact (if any) from Mongo
	var params map[string]string
	params = map[string]string{}

	if param.Publisher != nil {
		params["publisher"] = string(*param.Publisher)
	}

	if param.Purpose != nil {
		params["purpose"] = string(*param.Purpose)
	}

	if param.Name != nil {
		params["name"] = string(*param.Name)
	}

	if param.Language != nil {
		params["language"] = string(*param.Language)
	}

	questionnaireArr, err := FindRecords(&params, &model.Questionnaire{}, u.dal)
	if err != nil {
		return nil, err
	}

	var qArr []*model.Questionnaire
	for _, qr := range questionnaireArr {
		var q *model.Questionnaire
		bsonBytes, _ := bson.Marshal(qr)
		bson.Unmarshal(bsonBytes, &q)
		qArr = append(qArr, q)
	}

	return &qArr, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateQuestionnaire will create a new Questionnaire in Mongo using the Data Access Layer
func (u *QuestionnaireService) CreateQuestionnaire(questionnaire *model.Questionnaire) (*model.Questionnaire, error) {
	_, err := u.dal.Post(questionnaire)
	if err != nil {
		return nil, err
	}

	return questionnaire, nil
}
