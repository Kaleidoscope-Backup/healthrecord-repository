package service_test

import (
	"strconv"
	"time"

	"gitlab.com/karte/healthrecord-repository/model"
	"gitlab.com/karte/healthrecord-repository/util"
)

func createCompleteEMRHealthRecordObjectForTest(recordType model.HealthRecordType, prevRecord string) *model.HealthRecord {
	healthRecord := &model.HealthRecord{}

	//health record fields
	healthRecord.RecordType = recordType
	healthRecord.TransactionType = model.INSERT
	healthRecord.Name = string(recordType) + " test name"
	d := "Test Description"
	healthRecord.Description = &d
	healthRecord.Source = "EMR"
	healthRecord.ConsumerID = "12345"

	var now time.Time
	now = time.Now()
	t := util.Time{now}
	healthRecord.Occurred = t
	healthRecord.Created = t
	o := "1234567890"
	healthRecord.Organization = &o

	var sourceRecordID model.SourceRecordID
	sourceRecordID.System = "System"
	sourceRecordID.Value = "System Value"
	healthRecord.SourceRecordID = &sourceRecordID
	healthRecord.PreviousRecord = &prevRecord

	return healthRecord
}

func createTestDiagnosisArrayWithAllFields() *[]model.Diagnosis {
	var diagnosis []model.Diagnosis
	for i := 0; i < 2; i++ {
		diagnosis = append(diagnosis, createTestDiagnosisObjectWithAllFields(i))
	}
	return &diagnosis
}

func createTestDiagnosisObjectWithAllFields(index int) model.Diagnosis {
	diagnosis := model.Diagnosis{}
	diagnosis.Name = "TestDiagnosisName" + strconv.Itoa(index)
	//diagnosis.Code = createTestClinicalCodeObjectWithAllFields()
	//diagnosis.Code.SystemType = *createTestClinicalCodeTypeObjectWithAllFields()
	return diagnosis
}

func createTestReasonArrayWithAllFields() *[]model.Reason {
	var reasons []model.Reason
	for i := 0; i < 2; i++ {
		reasons = append(reasons, createTestReasonObjectWithAllFields(i))
	}
	return &reasons
}

func createTestReasonObjectWithAllFields(index int) model.Reason {
	reason := model.Reason{}
	reason.Name = "TestReasonName" + strconv.Itoa(index)
	//reason.Code = createTestClinicalCodeObjectWithAllFields()
	//reason.Code.SystemType = *createTestClinicalCodeTypeObjectWithAllFields()
	return reason
}
