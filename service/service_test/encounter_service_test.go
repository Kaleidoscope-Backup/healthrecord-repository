package service_test

import (
	"fmt"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/karte/healthrecord-repository/model"
	. "github.com/karte/healthrecord-repository/service"
	"github.com/karte/healthrecord-repository/util"
)

var _ = Describe("EncounterRecordService", func() {

	var (
		er        *model.EncounterRecord
		erService *EncounterRecordService
	)

	BeforeEach(func() {
		er = &model.EncounterRecord{}
		erService = NewEncounterRecordService(testDAL, testLog)
	})

	Describe("Validating creating a EncounterRecord in our MongoDB", func() {
		Context("With all fields populated in EncounterRecord", func() {
			var id string

			It("Should create a EncounterRecord without error and return an ID", func() {
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
				er.ConsumerID = "1234"

				//execute service
				er, err := erService.CreateEncounterRecord(er)
				if err != nil {
					fail := fmt.Sprintf("DB Create Failed: %v", err)
					Fail(fail)
				}
				id = er.Id
				Expect(er.Id).NotTo(Equal(nil))
			})
			It("Should be able to retrieve the new EncounterRecord by the returned ID", func() {
				v, _ := erService.FindById(id)
				Expect(v.RecordType).To(Equal(model.ENCOUNTER))

				var reasons []model.Reason
				reasons = *v.Reasons
				Expect(reasons[0].Name).To(Equal("TestReasonName0"))

				var diagnosis []model.Diagnosis
				diagnosis = *v.Diagnosis
				Expect(diagnosis[0].Name).To(Equal("TestDiagnosisName0"))

				var orders []model.EncounterOrder
				orders = *v.Orders
				Expect(orders[0].Name).To(Equal("TestEncounterOrderName0"))

				var prescriptions []model.Medication
				prescriptions = *v.Prescriptions
				Expect(prescriptions[0].ProductName).To(Equal("Tylenol"))

			})
		})

		Context("With missing required fields in EncounterRecord", func() {
			It("Should not create a EncounterRecord and Should Provide an Error", func() {
				_, err := erService.CreateEncounterRecord(er)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})

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

	//encounterOrder.Code = createTestClinicalCodeObjectWithAllFields()
	//encounterOrder.Code.SystemType = *createTestClinicalCodeTypeObjectWithAllFields()

	return encounterOrder
}
