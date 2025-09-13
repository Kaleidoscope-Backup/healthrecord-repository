package service

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	logging "github.com/op/go-logging"
	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/util"
	"gitlab.com/karte/mongo-lib/mserver"
)

/*==========================================================================================
QuestionnaireResponse service
==========================================================================================*/

//QuestionnaireResponseService is for creating QuestionnaireResponse
type QuestionnaireResponseService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

//NewQuestionnaireResponseService creates a new QuestionnaireResponse service
func NewQuestionnaireResponseService(dal mserver.DataAccessLayer, log *logging.Logger) *QuestionnaireResponseService {
	return &QuestionnaireResponseService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

//FindByID ..
func (u *QuestionnaireResponseService) FindByID(id string) (*model.QuestionnaireResponse, error) {

	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Contact (if any) from Mongo
	p, err := u.dal.Get(id, &model.QuestionnaireResponse{})
	if err != nil {
		return nil, err
	}

	var questionnaireResponse *model.QuestionnaireResponse
	bsonBytes, _ := bson.Marshal(p)
	bson.Unmarshal(bsonBytes, &questionnaireResponse)

	return questionnaireResponse, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

//CreateQuestionnaireResponse will create a new QuestionnaireResponse in Mongo using the Data Access Layer
func (u *QuestionnaireResponseService) CreateQuestionnaireResponse(questionnaireResponse *model.QuestionnaireResponse) (*model.QuestionnaireResponse, error) {
	_, err := u.dal.Post(questionnaireResponse)
	if err != nil {
		return nil, err
	}

	return questionnaireResponse, nil
}

//FindByConsumerID ..
func (u *QuestionnaireResponseService) FindByConsumerID(id string) (*[]*model.QuestionnaireResponse, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	var params map[string]string
	params = map[string]string{}

	params["consumerID"] = string(id)

	//find the matching questionnaire response Record (if any) from Mongo
	mrArr, err := FindRecords(&params, &model.QuestionnaireResponse{}, u.dal)
	if err != nil {
		return nil, err
	}

	var questionnaireResponseArr []*model.QuestionnaireResponse
	for _, ar := range mrArr {
		var questionnaireResponse *model.QuestionnaireResponse
		bsonBytes, _ := bson.Marshal(ar)
		bson.Unmarshal(bsonBytes, &questionnaireResponse)
		questionnaireResponseArr = append(questionnaireResponseArr, questionnaireResponse)
	}

	return &questionnaireResponseArr, nil
}

// Export ...
func (u *QuestionnaireResponseService) Export(id string) (*[]model.HealthRecordExportElement, error) {
	if id == "" {
		return nil, errors.New("Missing parameter")
	}

	questionnaireResponseRecords, _ := u.FindByConsumerID(id)
	if questionnaireResponseRecords != nil {
		questionnaireResponseRecordsIterable := *questionnaireResponseRecords
		records := []model.HealthRecordExportElement{}
		for _, questionnaireResponseRecord := range questionnaireResponseRecordsIterable {
			recordElements, _ := u.ExportElements(questionnaireResponseRecord)
			records = append(records, *recordElements...)
		}
		return &records, nil
	}

	return nil, nil
}

// ExportElements ...
func (u *QuestionnaireResponseService) ExportElements(questionnaireResponse *model.QuestionnaireResponse) (*[]model.HealthRecordExportElement, error) {
	if questionnaireResponse == nil {
		return nil, errors.New("Missing parameter")
	}

	// all records array
	records := &[]model.HealthRecordExportElement{}
	answers := questionnaireResponse.Items

	var err error
	records, err = u.ExportAnswers(questionnaireResponse.ConsumerID, &questionnaireResponse.TimeStamp, answers)
	if err != nil {
		return nil, err
	}

	return records, nil
}

// ExportAnswers ...
func (u *QuestionnaireResponseService) ExportAnswers(consumerID string, timeStamp *util.Time, answers *[]model.Answer) (*[]model.HealthRecordExportElement, error) {
	if answers == nil {
		return nil, errors.New("Missing parameter")
	}

	// Record ID for all options
	recordID := util.UUID()
	// all records array
	records := []model.HealthRecordExportElement{}

	for _, answer := range *answers {

		//Value Question
		valueQuestion := model.Value{}
		valueQuestion.ValueType = model.TEXT
		valueQuestion.ValueText = &answer.QuestionText

		// Populate record
		recordQuestion := model.HealthRecordExportElement{}
		recordQuestion.Name = "Answer"
		recordQuestion.TimeStamp = util.Time{timeStamp.Time}
		recordQuestion.Value = valueQuestion
		recordQuestion.RecordID = recordID
		recordQuestion.ConsumerID = consumerID

		//populate the record
		records = append(records, recordQuestion)

		// If Value is provided then
		if answer.AnswerValue != nil {
			//Value
			value := model.Value{}
			value = *answer.AnswerValue

			// Populate record
			record := model.HealthRecordExportElement{}
			record.Name = "Answer Value"
			record.TimeStamp = util.Time{timeStamp.Time}
			record.Value = value
			record.RecordID = recordID
			record.ConsumerID = consumerID

			//populate the record
			records = append(records, record)
		}

		// If selected option
		if answer.SelectedOptions != nil {
			selectedOptions := *answer.SelectedOptions
			for _, option := range selectedOptions {
				//Value
				value := model.Value{}
				value.ValueType = model.TEXT
				value.ValueText = &option.Option

				// Populate record
				record := model.HealthRecordExportElement{}
				record.Name = "Answer Option"
				record.TimeStamp = util.Time{timeStamp.Time}
				record.Value = value
				record.RecordID = recordID
				record.ConsumerID = consumerID

				//populate the record
				records = append(records, record)
			}
		}

		// Nested records ...
		if answer.Items != nil {
			nestedRecords, _ := u.ExportAnswers(consumerID, timeStamp, answer.Items)
			if nestedRecords != nil {
				records = append(records, *nestedRecords...)
			}
		}
	}

	return &records, nil
}
