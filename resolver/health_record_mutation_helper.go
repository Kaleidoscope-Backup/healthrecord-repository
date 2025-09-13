package resolver

import (
	"errors"
	"time"

	"github.com/karte/healthrecord-repository/constant"
	"github.com/karte/healthrecord-repository/model"
	"github.com/karte/healthrecord-repository/service"
	"github.com/karte/healthrecord-repository/util"
	logging "github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateSymptomFromInput ...
func CreateSymptomFromInput(ctx context.Context, input *model.SymptomInput) (*model.Symptom, error) {
	if input != nil {
		symptom := &model.Symptom{}
		symptom.Name = input.Name
		if input.Code != nil {
			code, err := CreateCodableConceptFromInput(ctx, input.Code)
			if err != nil {
				ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
				return nil, err
			}
			symptom.Code = code
		}
		return symptom, nil
	}
	return nil, nil
}

// CreateOnsetFromInput ...
func CreateOnsetFromInput(onsetInput *model.OnsetInput) *model.Onset {
	var onset *model.Onset
	onset = &model.Onset{}
	onset.Age = onsetInput.Age
	onset.Date = onsetInput.Date
	onset.Note = onsetInput.Note

	return onset
}

// CreateAbatementFromInput ...
func CreateAbatementFromInput(abatementInput *model.AbatementInput) *model.Abatement {
	var abatement *model.Abatement
	abatement = &model.Abatement{}
	abatement.Age = abatementInput.Age
	abatement.Date = abatementInput.Date
	abatement.Note = abatementInput.Note

	return abatement
}

// CreateHealthRecord ...
func CreateHealthRecord(ctx context.Context, healthRecordCreate *model.HealthRecordCreate, recordType model.HealthRecordType) (*model.HealthRecord, error) {
	healthRecord := &model.HealthRecord{}

	//check for consumer if the consumer does not exist throw error
	consumer, _ := ctx.Value(constant.ConsumerService).(*service.ConsumerService).FindByID(healthRecordCreate.ConsumerID)

	if consumer == nil {
		return nil, errors.New("Invalid consumer ID")
	}

	//health record fields
	healthRecord.RecordType = recordType
	healthRecord.TransactionType = model.INSERT
	healthRecord.Name = healthRecordCreate.Name
	healthRecord.Description = healthRecordCreate.Description
	healthRecord.Source = healthRecordCreate.Source
	healthRecord.ConsumerID = healthRecordCreate.ConsumerID

	// Geo location
	if healthRecordCreate.Location != nil {
		location := CreateGeolocationFromInput(healthRecordCreate.Location)
		healthRecord.Location = location
	}

	var now time.Time
	now = time.Now()
	t := util.Time{now}
	healthRecord.Occurred = healthRecordCreate.Occurred
	healthRecord.Created = t
	healthRecord.Organization = healthRecordCreate.Organization

	if healthRecordCreate.SourceRecordID != nil {
		var sourceRecordID model.SourceRecordID
		sourceRecordID.System = healthRecordCreate.SourceRecordID.System
		sourceRecordID.Value = healthRecordCreate.SourceRecordID.Value
		healthRecord.SourceRecordID = &sourceRecordID
	}

	if healthRecordCreate.References != nil && len(*healthRecordCreate.References) > 0 {

		var referenceInputArray []model.ReferenceHealthRecordInput
		referenceInputArray = *healthRecordCreate.References

		var references []model.ReferenceHealthRecord
		references = []model.ReferenceHealthRecord{}

		for i := 0; i < len(*healthRecordCreate.References); i++ {
			var reference model.ReferenceHealthRecord
			referenceInput := referenceInputArray[i]
			reference.Type = referenceInput.Type
			reference.ReferencedID = referenceInput.ReferencedID
			references = append(references, reference)
		}

		healthRecord.References = &references
	}

	return healthRecord, nil
}

// CreateReasons will create an array of Reasons object from an Array of ReasonCreate input objects
func createReasons(reasonsCreate *[]model.ReasonInput) *[]model.Reason {
	var reasons []model.Reason

	if reasonsCreate != nil && len(*reasonsCreate) > 0 {
		for i := 0; i < len(*reasonsCreate); i++ {
			var rc model.ReasonInput
			rc = (*reasonsCreate)[i]
			if r := createReason(&rc); r != nil {
				reasons = append(reasons, *r)
			}
		}
		return &reasons
	}

	return nil
}

// createReason will create a medication object from a MedicationCreate input object
func createReason(reasonCreate *model.ReasonInput) *model.Reason {
	reason := &model.Reason{}

	reason.Name = reasonCreate.Name

	//resolve matching clinical code
	c := resolveClinicalCodeFromName(reasonCreate.Name)
	if c != nil {
		reason.Code = c
	}

	return reason
}

// CreateDiagnosisArray will create an array of Diagnosis objects from an Array of DiagnosisCreate input objects
func createDiagnosisArray(diagnosisCreateArray *[]model.DiagnosisInput) *[]model.Diagnosis {
	var diagnosisArray []model.Diagnosis

	if diagnosisCreateArray != nil && len(*diagnosisCreateArray) > 0 {
		for i := 0; i < len(*diagnosisCreateArray); i++ {
			var dc model.DiagnosisInput
			dc = (*diagnosisCreateArray)[i]
			if r := createDiagnosis(&dc); r != nil {
				diagnosisArray = append(diagnosisArray, *r)
			}
		}
		return &diagnosisArray
	}

	return nil
}

// createReason will create a medication object from a MedicationCreate input object
func createDiagnosis(diagnosisCreate *model.DiagnosisInput) *model.Diagnosis {
	diagnosis := &model.Diagnosis{}

	diagnosis.Name = diagnosisCreate.Name

	//resolve matching clinical code
	c := resolveClinicalCodeFromName(diagnosisCreate.Name)
	if c != nil {
		diagnosis.Code = c
	}

	return diagnosis
}
