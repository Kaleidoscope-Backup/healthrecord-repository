package service

import (
	"errors"
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/util"
	"github.com/karte/mongo-lib/mserver"
	logging "github.com/op/go-logging"
)

/*==========================================================================================
Procedure Record service
==========================================================================================*/

// ProcedureRecordService is for creating habit
type ProcedureRecordService struct {
	dal mserver.DataAccessLayer
	log *logging.Logger
}

// NewProcedureRecordService creates a new Procedure service that has all calls to the database, queries and mutations via the Data Access Layer
func NewProcedureRecordService(dal mserver.DataAccessLayer, log *logging.Logger) *ProcedureRecordService {
	return &ProcedureRecordService{dal: dal, log: log}
}

/*==========================================================================================
Query Operations
==========================================================================================*/

// FindByID ..
func (u *ProcedureRecordService) FindByID(id string) (*model.ProcedureRecord, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching Procedure Record (if any) from Mongo
	p, err := u.dal.Get(id, &model.ProcedureRecord{})
	if err != nil {
		return nil, err
	}

	var procedure *model.ProcedureRecord
	bsonBytes, _ := bson.Marshal(p)
	bson.Unmarshal(bsonBytes, &procedure)

	return procedure, nil
}

// FindByConsumerID ..
func (u *ProcedureRecordService) FindByConsumerID(id string) (*[]*model.ProcedureRecord, error) {
	if id == "" {
		return nil, errors.New("Missing parameter id")
	}

	//find the matching ProcedureRecord (if any) from Mongo
	prArr, err := FindHealthRecordsByConsumerID(id, &model.ProcedureRecord{}, u.dal)
	if err != nil {
		return nil, err
	}

	var procedureRecordArr []*model.ProcedureRecord
	for _, pr := range prArr {
		var procedureRecord *model.ProcedureRecord
		bsonBytes, _ := bson.Marshal(pr)
		bson.Unmarshal(bsonBytes, &procedureRecord)
		procedureRecordArr = append(procedureRecordArr, procedureRecord)
	}

	return &procedureRecordArr, nil
}

/*==========================================================================================
Mutation Operations
==========================================================================================*/

// CreateProcedureRecord will create a new procedure in Mongo using the Data Access Layer ...
func (u *ProcedureRecordService) CreateProcedureRecord(procedureRecord *model.ProcedureRecord) (*model.ProcedureRecord, error) {

	if &procedureRecord.Category == nil || &procedureRecord.Status == nil {
		return nil, errors.New("Missing a required field aborting before saving to the DB")
	}

	hError := ValidateHealthRecord(&procedureRecord.HealthRecord)
	if hError != nil {
		fmt.Printf("DB Create Failed For Addiction Record Error : %v", hError)
		return nil, hError
	}

	id, err := u.dal.Post(procedureRecord)
	if err != nil {
		fmt.Printf("DB Create Failed For Addiction Record Error : %v", err)
		return nil, err
	}

	errHr := u.dal.PostWithID(id, &procedureRecord.HealthRecord)
	if errHr != nil {
		return nil, err
	}

	PostRecord(procedureRecord, procedureRecord.ConsumerID, id)
	return procedureRecord, nil
}

// Export ...
func (u *ProcedureRecordService) Export(id string) (*[]model.HealthRecordExportElement, error) {
	if id == "" {
		return nil, errors.New("Missing parameter")
	}

	procedureRecords, _ := u.FindByConsumerID(id)
	if procedureRecords != nil {
		procedureRecordsIterable := *procedureRecords
		records := []model.HealthRecordExportElement{}
		for _, procedureRecord := range procedureRecordsIterable {
			recordElements, _ := u.ExportElements(procedureRecord)
			records = append(records, *recordElements...)
		}
		return &records, nil
	}

	return nil, nil
}

// ExportElements ...
func (u *ProcedureRecordService) ExportElements(procedure *model.ProcedureRecord) (*[]model.HealthRecordExportElement, error) {
	if procedure == nil {
		return nil, errors.New("Missing parameter")
	}

	//Location
	location := procedure.Location

	// all records array
	records := []model.HealthRecordExportElement{}
	recordID := util.UUID()

	// Category
	valueCategory := model.Value{}
	valueCategory.ValueType = model.TEXT
	categoryText := string(procedure.Category)
	valueCategory.ValueText = &categoryText

	// Record
	recordCategory := model.HealthRecordExportElement{}
	recordCategory.Name = "Procedure Category"
	recordCategory.TimeStamp = procedure.Occurred
	recordCategory.ConsumerID = procedure.ConsumerID
	recordCategory.RecordID = recordID
	recordCategory.Value = valueCategory
	recordCategory.Location = location

	//populate the record
	records = append(records, recordCategory)

	// Reason
	valueReason := model.Value{}
	valueReason.ValueType = model.TEXT
	valueReason.ValueText = &procedure.Reason

	// Record
	recordReason := model.HealthRecordExportElement{}
	recordReason.Name = "Procedure Reason"
	recordReason.TimeStamp = procedure.Occurred
	recordReason.ConsumerID = procedure.ConsumerID
	recordReason.RecordID = recordID
	recordReason.Value = valueReason
	recordReason.Location = location

	//populate the record
	records = append(records, recordReason)

	if procedure.Performer != nil {
		value := model.Value{}
		value.ValueType = model.TEXT
		value.ValueText = procedure.Performer

		// Record
		record := model.HealthRecordExportElement{}
		record.Name = "Procedure Performer"
		record.TimeStamp = procedure.Occurred
		record.ConsumerID = procedure.ConsumerID
		record.RecordID = recordID
		record.Value = value
		record.Location = location

		//populate the record
		records = append(records, record)
	}

	if procedure.BodySite != nil {
		value := model.Value{}
		value.ValueType = model.TEXT
		value.ValueText = procedure.BodySite

		// Record
		record := model.HealthRecordExportElement{}
		record.Name = "Procedure Body Site"
		record.TimeStamp = procedure.Occurred
		record.ConsumerID = procedure.ConsumerID
		record.RecordID = recordID
		record.Value = value
		record.Location = location

		//populate the record
		records = append(records, record)
	}

	if procedure.Outcome != nil {
		value := model.Value{}
		value.ValueType = model.TEXT
		outcome := string(*procedure.Outcome)
		value.ValueText = &outcome

		// Record
		record := model.HealthRecordExportElement{}
		record.Name = "Procedure Outcome"
		record.TimeStamp = procedure.Occurred
		record.ConsumerID = procedure.ConsumerID
		record.RecordID = recordID
		record.Value = value
		record.Location = location

		//populate the record
		records = append(records, record)
	}

	if procedure.FollowupInstruction != nil {
		value := model.Value{}
		value.ValueType = model.TEXT
		value.ValueText = procedure.FollowupInstruction

		// Record
		record := model.HealthRecordExportElement{}
		record.Name = "Procedure Followup Instructions"
		record.TimeStamp = procedure.Occurred
		record.ConsumerID = procedure.ConsumerID
		record.RecordID = recordID
		record.Value = value
		record.Location = location

		//populate the record
		records = append(records, record)
	}

	if procedure.Report != nil {
		value := model.Value{}
		value.ValueType = model.TEXT
		value.ValueText = procedure.Report

		// Record
		record := model.HealthRecordExportElement{}
		record.Name = "Procedure Report"
		record.TimeStamp = procedure.Occurred
		record.ConsumerID = procedure.ConsumerID
		record.RecordID = recordID
		record.Value = value
		record.Location = location

		//populate the record
		records = append(records, record)
	}

	return &records, nil
}
