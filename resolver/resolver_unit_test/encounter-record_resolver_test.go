package resolver_unit_test

import (
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/resolver"
	"github.com/karte/healthrecord-repository/util"
)

var _ = Describe("EncounterRecordResolver", func() {

	var (
		er *model.EncounterRecord
	)

	BeforeEach(func() {
		er = &model.EncounterRecord{}
	})

	Describe("Validating Encounter Record Resolver", func() {
		Context("With All Fields Requested", func() {
			It("Should resolve values for all encounter record fields", func() {
				er = createTestEncounterRecordObjectWithAllFields()
				hr := HealthRecordResolver{&er.HealthRecord}
				encounterRecordResolver := EncounterRecordResolver{hr, er}

				Expect(encounterRecordResolver.RecordType()).To(Equal(er.RecordType))
				Expect(encounterRecordResolver.AttendedBy().FirstName()).To(Equal(er.AttendedBy.FirstName))
				Expect((*encounterRecordResolver.Reasons())[0].Name()).To(Equal((*er.Reasons)[0].Name))
				Expect((*encounterRecordResolver.Diagnosis())[0].Name()).To(Equal((*er.Diagnosis)[0].Name))
				Expect((*encounterRecordResolver.Orders())[0].ProcedureCode()).To(Equal((*er.Orders)[0].ProcedureCode))
				Expect((*encounterRecordResolver.Prescriptions())[0].ProductName()).To(Equal((*er.Prescriptions)[0].ProductName))
			})
		})
	})
})

func createTestEncounterRecordObjectWithAllFields() *model.EncounterRecord {
	er := &model.EncounterRecord{}

	//health record
	previousRecord := "7890-890"
	if hr := createCompleteEMRHealthRecordObjectForTest(model.ENCOUNTER, previousRecord); hr != nil {
		er.HealthRecord = *hr
	}

	//encounter
	er.AttendedBy = createTestPractitionerObjectWithAllFields()
	er.Reasons = createTestReasonArrayWithAllFields()
	er.Diagnosis = createTestDiagnosisArrayWithAllFields()
	er.Prescriptions = createTestMedicationsArrayWithAllFields()
	er.Orders = createTestOrdersArray()

	return er
}

func createTestOrdersArray() *[]model.EncounterOrder {
	var orders []model.EncounterOrder

	for i := 0; i < 2; i++ {
		orders = append(orders, *createTestEncounterOrderObjectWithAllFields(i))
	}

	return &orders
}

func createTestEncounterOrderObjectWithAllFields(index int) (encounterOrder *model.EncounterOrder) {
	encounterOrder = &model.EncounterOrder{}
	encounterOrder.Name = "TestEncounterOrderName" + strconv.Itoa(index)
	encounterOrder.ProcedureCode = model.CODE_103693007 // Diagnostic Procedure
	tt := "TestType"
	encounterOrder.Type = &tt

	var now time.Time
	now = time.Now()
	t := util.Time{now}
	encounterOrder.ExpectedDate = &t
	encounterOrder.ExpirationDate = &t

	encounterOrder.Code = createTestClinicalCodeObjectWithAllFields()
	return encounterOrder
}

func createTestClinicalCodeObjectWithAllFields() (clinicalCode *model.ClinicalCode) {
	//Clinical Code Fields
	clinicalCode = &model.ClinicalCode{}
	clinicalCode.Code = "TestCode"
	return clinicalCode
}

func createCompleteEMRHealthRecordObjectForTest(recordType model.HealthRecordType, prevRecord string) *model.HealthRecord {
	healthRecord := &model.HealthRecord{}

	//health record fields
	healthRecord.RecordType = recordType
	healthRecord.TransactionType = model.INSERT
	healthRecord.Name = string(recordType) + " test name"
	d := "Test Description"
	healthRecord.Description = &d
	healthRecord.Source = "EMR"

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
	diagnosis.Code = createTestClinicalCodeObjectWithAllFields()
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
	reason.Code = createTestClinicalCodeObjectWithAllFields()
	return reason
}

func createTestPractitionerObjectWithAllFields() *model.Practitioner {
	practitioner := &model.Practitioner{}

	practitioner = createTestPractitionerObjectWithRequiredFields()
	lp := "English"
	practitioner.LanguagePreference = &lp //non-required fields are pointers

	return practitioner
}

func createTestPractitionerObjectWithRequiredFields() *model.Practitioner {
	practitioner := &model.Practitioner{}

	practitioner.FirstName = "testFirstName"
	practitioner.LastName = "testLastName"
	practitioner.Email = "test@gmail.com"

	practitioner.Speciality = "Nuerosurgeon"
	practitioner.Qualification = "MD"
	practitioner.Organization = "Massachussets General Hospital"

	return practitioner
}

func createTestMedicationsArrayWithAllFields() *[]model.Medication {
	var medications []model.Medication

	medications = append(medications, *createTestMedicationObjectWithAllFields("Tylenol"))
	medications = append(medications, *createTestMedicationObjectWithAllFields("Sudafed"))

	return &medications
}

func createTestMedicationObjectWithAllFields(productName string) *model.Medication {
	medication := &model.Medication{}

	medication.MedicationStatus = "ACTIVE"
	medication.ProductName = productName
	medication.IsOverTheCounter = true
	medication.Route = model.ORAL_ADMINISTRATION
	medication.Instructions = "Take with plenty of water"

	var dosage model.Dosage
	medication.Dosage = &dosage
	medication.Dosage.Frequency = "2 times a day"
	medication.Dosage.Unit = "5 mg"
	medication.Dosage.Value = 1

	now := time.Now()
	var tr = util.Time{now}
	medication.Start = tr

	return medication
}
